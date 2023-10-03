package webhandler

type CustomError struct {
	error string `json:"error"`
}

func CustomErrorResponse(msg string) *CustomError {
	return &CustomError{
		error: msg,
	}
}
