package main

import (
	"context"
	"errors"
	"strings"
)

// ErrEmpty creates an error if the string is empty
var ErrEmpty = errors.New("Empty String")

// StringService is an interface for core business logic
type StringService interface {
	Uppercase(context.Context, string) (string, error)
	Count(context.Context, string) int
}

type stringService struct{}

func (stringService) Uppercase(_ context.Context, s string) (string, error) {
	if s == " " {
		return " ", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (stringService) Count(_ context.Context, s string) int {
	return len(s)
}
