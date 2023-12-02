package chat

import (
	"astroboy/internal/dependencies"
	"astroboy/internal/model"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"time"
)

type HistoryService struct {
	deps *dependencies.Dependencies
}

func NewHistoryService(deps *dependencies.Dependencies) *HistoryService {
	return &HistoryService{deps: deps}
}

func (hs *HistoryService) ForUser(userId string) ([]model.ChatMessage, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 1*time.Second)
	defer cancel()

	keyEx := expression.Key("user_id").Equal(expression.Value(userId))
	expr, err := expression.NewBuilder().WithKeyCondition(keyEx).Build()
	if err != nil {
		return nil, err
	}

	response, err := hs.deps.Db.Client.Query(ctx, &dynamodb.QueryInput{
		TableName:                 aws.String(hs.deps.Db.TableName),
		IndexName:                 aws.String(hs.deps.Db.UserIndexName),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
	})

	var chatMessages []model.ChatMessage
	err = attributevalue.UnmarshalListOfMaps(response.Items, &chatMessages)
	if err != nil {
		return nil, err
	}

	return chatMessages, nil
}
