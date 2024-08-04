package errors

import "net/http"

type AppErr interface {
	GetMessage() string
	GetStatus() int
}

type appErr struct {
	message string
	status  int
}

func (ap *appErr) GetMessage() string {
	return ap.message
}

func (ap *appErr) GetStatus() int {
	return ap.status
}

func NewAppErr(message string, status int) AppErr {
	return &appErr{
		message: message,
		status:  status,
	}
}

func NewBadRequestErr(message string) AppErr {
	return &appErr{
		message: message,
		status:  http.StatusBadRequest,
	}
}

func NewUnauthorizedErr() AppErr {
	return &appErr{
		message: "Unauthorized action",
		status:  http.StatusUnauthorized,
	}
}

func NewInternalServerErr() AppErr {
	return &appErr{
		message: "Internal server error",
		status:  http.StatusInternalServerError,
	}
}
