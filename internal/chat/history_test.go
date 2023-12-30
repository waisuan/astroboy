package chat

import (
	"astroboy/internal/dependencies"
	"astroboy/internal/mocks"
	"astroboy/internal/model"
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/golang/mock/gomock"
	testify "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHistoryService_ForUser(t *testing.T) {
	assert := testify.New(t)
	must := require.New(t)

	t.Run("no chat history", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockDb := mocks.NewMockIDatabase(mockCtrl)
		mockDb.EXPECT().QueryWithIndex(gomock.Any(), dependencies.UserGsiName, gomock.Any()).Return(nil, nil)

		h := NewHistoryService(&dependencies.Dependencies{DB: mockDb})

		out, err := h.ForUser("esia")
		assert.Nil(err)
		assert.Empty(out)
	})

	t.Run("chat history exists", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockDb := mocks.NewMockIDatabase(mockCtrl)
		mockDb.EXPECT().QueryWithIndex(gomock.Any(), dependencies.UserGsiName, gomock.Any()).Return(dependencies.DbQueryOutput{
			{
				"id":        &types.AttributeValueMemberS{Value: "test123"},
				"timestamp": &types.AttributeValueMemberN{Value: "0"},
			},
		}, nil)

		h := NewHistoryService(&dependencies.Dependencies{DB: mockDb})

		out, err := h.ForUser("esia")
		assert.Nil(err)
		must.NotEmpty(out)
		assert.Equal("test123", out[0].Id)
		assert.Equal(int64(0), out[0].Timestamp)
	})

	t.Run("runtime error", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockDb := mocks.NewMockIDatabase(mockCtrl)
		mockDb.EXPECT().QueryWithIndex(gomock.Any(), dependencies.UserGsiName, gomock.Any()).Return(nil, errors.New("something bad happened"))

		h := NewHistoryService(&dependencies.Dependencies{DB: mockDb})

		out, err := h.ForUser("esia")
		must.NotNil(err)
		assert.Empty(out)
	})
}

func TestHistoryService_AddChatMessage(t *testing.T) {
	must := require.New(t)

	t.Run("successful", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockDb := mocks.NewMockIDatabase(mockCtrl)
		mockDb.EXPECT().PutItem(context.TODO(), gomock.AssignableToTypeOf(&model.ChatMessage{}), nil).Return(nil)

		h := NewHistoryService(&dependencies.Dependencies{DB: mockDb})
		err := h.AddChatMessage("esia", &model.ChatMessage{Body: "testing"})
		must.Nil(err)
	})

	t.Run("runtime error", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockDb := mocks.NewMockIDatabase(mockCtrl)
		mockDb.EXPECT().PutItem(context.TODO(), gomock.Any(), nil).Return(errors.New("something bad happened"))

		h := NewHistoryService(&dependencies.Dependencies{DB: mockDb})
		err := h.AddChatMessage("esia", &model.ChatMessage{})
		must.NotNil(err)
	})
}
