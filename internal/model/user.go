package model

type User struct {
	Id           string `dynamodbav:"id" json:"id"`
	Timestamp    int64  `dynamodbav:"timestamp" json:"timestamp"`
	Password     string `dynamodbav:"password" json:"-"`
	Email        string `dynamodbav:"email" json:"email"`
	RegisteredAt int64  `dynamodbav:"registered_at" json:"registered_at"`
}
