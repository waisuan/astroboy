package dependencies

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/pkg/errors"
	"log"
	"time"
)

const (
	UserGsiName  = "USER_GSI"
	ConvoGsiName = "CONVO_GSI"
)

type DB struct {
	Client         *dynamodb.Client
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
	}
}

// <message_id> | <timestamp> | <user_id> | <convo_id> | <body>
// - gsi: [user_id, timestamp]
// - gsi: [convo_id, timestamp]
func createTable(svc *dynamodb.Client, tableName string) error {
	o, _ := svc.DescribeTable(ctx, &dynamodb.DescribeTableInput{TableName: aws.String(tableName)})
	if o != nil && o.Table.TableStatus == types.TableStatusActive {
		_, err := svc.DeleteTable(ctx, &dynamodb.DeleteTableInput{TableName: aws.String(tableName)})
		if err != nil {
			return errors.Wrap(err, "failed to drop table")
		}
	}

	_, err := svc.CreateTable(ctx, &dynamodb.CreateTableInput{
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("message_id"),
				AttributeType: types.ScalarAttributeTypeS,
			},
			{
				AttributeName: aws.String("created_at"),
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
				AttributeName: aws.String("message_id"),
				KeyType:       types.KeyTypeHash,
			},
			{
				AttributeName: aws.String("created_at"),
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
						AttributeName: aws.String("created_at"),
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
						AttributeName: aws.String("created_at"),
						KeyType:       types.KeyTypeRange,
					},
				},
				Projection: &types.Projection{ProjectionType: types.ProjectionTypeAll},
			},
		},
		TableName:   aws.String(tableName),
		BillingMode: types.BillingModePayPerRequest,
	})
	//, func(options *dynamodb.Options) {
	//	options.Retryer = retry.AddWithErrorCodes(options.Retryer, (*types.ResourceInUseException)(nil).ErrorCode())
	//	options.Retryer = retry.AddWithMaxAttempts(options.Retryer, 0)
	//}

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
