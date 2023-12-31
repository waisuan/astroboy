package chat

import (
	"astroboy/internal/dependencies"
	"astroboy/internal/model"
	"context"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/google/uuid"
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

	keyEx := expression.Key(dependencies.UserGsiPKey).Equal(expression.Value(userId))
	expr, err := expression.NewBuilder().WithKeyCondition(keyEx).Build()
	if err != nil {
		return nil, err
	}

	out, err := hs.deps.DB.Query(ctx, expr, dependencies.UserGsiName)
	if err != nil {
		return nil, err
	}

	var chatMessages []model.ChatMessage
	err = attributevalue.UnmarshalListOfMaps(out, &chatMessages)
	if err != nil {
		return nil, err
	}

	return chatMessages, nil
}

func (hs *HistoryService) AddChatMessage(userId string, chatMsg *model.ChatMessage) error {
	chatMsg.Id = uuid.NewString()
	chatMsg.Timestamp = time.Now().UnixNano()
	chatMsg.UserId = userId
	chatMsg.ConvoId = uuid.NewString()

	return hs.deps.DB.PutItem(context.TODO(), chatMsg, nil)
}
