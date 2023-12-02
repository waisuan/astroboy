package model

type ChatMessage struct {
	MessageId string `dynamodbav:"message_id" json:"message_id"`
	CreatedAt int64  `dynamodbav:"created_at" json:"created_at"`
	Body      string `dynamodbav:"body" json:"body"`
	UserId    string `dynamodbav:"user_id" json:"user_id"`
	ConvoId   string `dynamodbav:"convo_id" json:"convo_id"`
}
