package chat

import (
	"astroboy/internal/dependencies"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"log"
)

type HistoryService struct {
	deps *dependencies.Dependencies
}

func NewHistoryService(deps *dependencies.Dependencies) *HistoryService {
	return &HistoryService{deps: deps}
}

func (hs *HistoryService) ForUser(userId string) error {
	out, err := hs.deps.Db.Client.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:              aws.String(hs.deps.Db.TableName),
		IndexName:              aws.String(hs.deps.Db.UserIndexName),
		KeyConditionExpression: aws.String("user_id = :userId"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":userId": &types.AttributeValueMemberS{Value: userId},
		},
	})
	if err != nil {
		return err
	}

	log.Printf("Result: %v\n", out.Items)

	return nil
}

//<message_id> | <timestamp> | <user_id> | <convo_id> | <body>
//- gsi: [user_id, timestamp]
//- gsi: [convo_id, timestamp]
