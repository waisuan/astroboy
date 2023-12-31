package dependencies

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/pkg/errors"
	"log"
	"time"
)

const (
	UserGsiName  = "USER_GSI"
	ConvoGsiName = "CONVO_GSI"
	UserGsiPKey  = "user_id"
	PartitionKey = "id"
)

type IDatabase interface {
	Query(ctx context.Context, expr expression.Expression, indexName string) (DbQueryOutput, error)
	PutItem(ctx context.Context, input interface{}, expr *expression.Expression) error
	ClearTable(ctx context.Context) error
}

type DbQueryOutput []map[string]types.AttributeValue

type DB struct {
	Client         *dynamodb.Client
	Config         *Config
	TableName      string
	UserIndexName  string
	ConvoIndexName string
}

func InitDB(cfg *Config) *DB {
	c, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("unable to init DB: %e\n", err)
	}

	svc := dynamodb.NewFromConfig(c, func(o *dynamodb.Options) {
		if cfg.DevMode {
			o.BaseEndpoint = aws.String(fmt.Sprintf("http://%s:%s", cfg.DbHost, cfg.DbPort))
		}
	})

	if cfg.DevMode {
		err := createTable(svc, cfg.DbTableName)
		if err != nil {
			panic(err)
		}

		log.Printf("Created table: %s\n", cfg.DbTableName)
	}

	return &DB{
		Client:         svc,
		TableName:      cfg.DbTableName,
		UserIndexName:  UserGsiName,
		ConvoIndexName: ConvoGsiName,
		Config:         cfg,
	}
}

// <message_id> | <timestamp> | <user_id> | <convo_id> | <body>
// - gsi: [user_id, timestamp]
// - gsi: [convo_id, timestamp]
func createTable(svc *dynamodb.Client, tableName string) error {
	o, _ := svc.DescribeTable(ctx, &dynamodb.DescribeTableInput{TableName: aws.String(tableName)})
	if o != nil && o.Table.TableStatus == types.TableStatusActive {
		return nil
	}

	_, err := svc.CreateTable(ctx, &dynamodb.CreateTableInput{
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("id"),
				AttributeType: types.ScalarAttributeTypeS,
			},
			{
				AttributeName: aws.String("timestamp"),
				AttributeType: types.ScalarAttributeTypeN,
			},
			{
				AttributeName: aws.String("user_id"),
				AttributeType: types.ScalarAttributeTypeS,
			},
			{
				AttributeName: aws.String("convo_id"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("id"),
				KeyType:       types.KeyTypeHash,
			},
			{
				AttributeName: aws.String("timestamp"),
				KeyType:       types.KeyTypeRange,
			},
		},
		GlobalSecondaryIndexes: []types.GlobalSecondaryIndex{
			{
				IndexName: aws.String(UserGsiName),
				KeySchema: []types.KeySchemaElement{
					{
						AttributeName: aws.String("user_id"),
						KeyType:       types.KeyTypeHash,
					},
					{
						AttributeName: aws.String("timestamp"),
						KeyType:       types.KeyTypeRange,
					},
				},
				Projection: &types.Projection{ProjectionType: types.ProjectionTypeAll},
			},
			{
				IndexName: aws.String(ConvoGsiName),
				KeySchema: []types.KeySchemaElement{
					{
						AttributeName: aws.String("convo_id"),
						KeyType:       types.KeyTypeHash,
					},
					{
						AttributeName: aws.String("timestamp"),
						KeyType:       types.KeyTypeRange,
					},
				},
				Projection: &types.Projection{ProjectionType: types.ProjectionTypeAll},
			},
		},
		TableName:   aws.String(tableName),
		BillingMode: types.BillingModePayPerRequest,
	})

	if err != nil {
		return errors.Wrap(err, "failed to create table")
	}

	w := dynamodb.NewTableExistsWaiter(svc)
	err = w.Wait(
		ctx,
		&dynamodb.DescribeTableInput{
			TableName: aws.String(tableName),
		},
		2*time.Minute,
		func(o *dynamodb.TableExistsWaiterOptions) {
			o.MaxDelay = 5 * time.Second
			o.MinDelay = 5 * time.Second
		},
	)
	if err != nil {
		return errors.Wrap(err, "timed out while waiting for table to become active")
	}

	return nil
}

func (db *DB) ClearTable(ctx context.Context) error {
	if !db.Config.DevMode {
		return errors.New("unable to run command when in non-dev mode")
	}

	_, err := db.Client.DeleteTable(ctx, &dynamodb.DeleteTableInput{TableName: aws.String(db.TableName)})
	if err != nil {
		return errors.Wrap(err, "failed to delete table")
	}

	w := dynamodb.NewTableNotExistsWaiter(db.Client)
	err = w.Wait(
		ctx,
		&dynamodb.DescribeTableInput{
			TableName: aws.String(db.TableName),
		},
		2*time.Minute,
		func(o *dynamodb.TableNotExistsWaiterOptions) {
			o.MaxDelay = 5 * time.Second
			o.MinDelay = 5 * time.Second
		},
	)
	if err != nil {
		return errors.Wrap(err, "timed out while waiting for table to become inactive")
	}

	err = createTable(db.Client, db.TableName)
	if err != nil {
		return errors.Wrap(err, "failed to recreate table")
	}

	return nil
}

func (db *DB) Query(ctx context.Context, expr expression.Expression, indexName string) (DbQueryOutput, error) {
	queryInput := &dynamodb.QueryInput{
		TableName:                 aws.String(db.TableName),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
	}

	if indexName != "" {
		queryInput.IndexName = aws.String(indexName)
	}

	response, err := db.Client.Query(ctx, queryInput)
	if err != nil {
		return nil, err
	}

	return response.Items, nil
}

func (db *DB) PutItem(ctx context.Context, input interface{}, expr *expression.Expression) error {
	item, err := attributevalue.MarshalMap(input)
	if err != nil {
		return err
	}

	queryInput := &dynamodb.PutItemInput{
		TableName: aws.String(db.TableName),
		Item:      item,
	}

	if expr != nil {
		queryInput.ConditionExpression = expr.Condition()
		queryInput.ExpressionAttributeNames = expr.Names()
		queryInput.ExpressionAttributeValues = expr.Values()
	}

	_, err = db.Client.PutItem(ctx, queryInput)

	return err
}
