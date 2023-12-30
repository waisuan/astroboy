package model

type ChatMessage struct {
	Id        string `dynamodbav:"id" json:"id"`
	Timestamp int64  `dynamodbav:"timestamp" json:"timestamp"`
	Body      string `dynamodbav:"body" json:"body"`
	UserId    string `dynamodbav:"user_id" json:"user_id"`
	ConvoId   string `dynamodbav:"convo_id" json:"convo_id"`
}
