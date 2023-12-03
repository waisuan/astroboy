package chat

import (
	"astroboy/internal/dependencies"
	"github.com/golang/mock/gomock"
	testify "github.com/stretchr/testify/assert"
	"testing"
)

func TestHistoryService_ForUser(t *testing.T) {
	assert := testify.New(t)

	t.Run("no chat history", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockDb := dependencies.NewMockIDatabase(mockCtrl)
		mockDb.EXPECT().QueryWithIndex(gomock.Any(), dependencies.UserGsiName, gomock.Any()).Return(nil, nil)

		h := NewHistoryService(&dependencies.Dependencies{Db: mockDb})

		out, err := h.ForUser("esia")
		assert.Nil(err)
		assert.Empty(out)
	})
}
