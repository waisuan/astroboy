package model

type User struct {
	Username    string `json:"username" param:"username"`
	Email       string `json:"email"`
	DateOfBirth string `json:"date_of_birth"`
}
