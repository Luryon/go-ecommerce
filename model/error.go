package model

import "fmt"

type Error struct {
	Code       string
	Err        error
	Who        string
	StatusHttp int
	Data       interface{}
	APIMessage string
	UserID     string
}

func NewError() Error {
	return Error{}
}

func (e *Error) Error() string {
	return fmt.Sprintf("Error code: %s, Error: %v, Who: %s, Status: %d, Data: %v, UserID: %s",
		e.Code,
		e.Err,
		e.Who,
		e.StatusHttp,
		e.Data,
		e.UserID)
}

func (e *Error) HasCode() bool {
	return e.Code != ""
}

func (e *Error) HasStatus() bool {
	return e.StatusHttp > 0
}

func (e *Error) HasData() bool {
	return e.Data != nil
}

func (e *Error) HasUserID() bool {
	return e.UserID != ""
}
