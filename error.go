package main

import "errors"

type Error struct {
	Code  int
	Error error
}

func NewError(code int, msg string) *Error {
	return &Error{Code: code, Error: errors.New(msg)}
}

var (
	systemError = NewError(500, "system error.")
	notFound    = NewError(404, "not found.")
)
