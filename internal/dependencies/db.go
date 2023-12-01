package dependencies

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
	"log"
)

type DB struct {
	Client         *dynamodb.Client
	TableName      string
	UserIndexName  string
	ConvoIndexName string
}

func InitDB(cfg *Config) *DB {
	c, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to init DB: %e\n", err)
	}

	svc := dynamodb.NewFromConfig(c, func(o *dynamodb.Options) {
		if cfg.DevMode {
			o.BaseEndpoint = aws.String("http://localhost:8080")
		}
	})

	return &DB{
		Client:         svc,
		TableName:      cfg.DbTableName,
		UserIndexName:  "USER_GSI",
		ConvoIndexName: "CONVO_GSI",
	}
}
