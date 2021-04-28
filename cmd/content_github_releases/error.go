package main

import (
	"errors"
	"strings"
)

var (
	ErrMissingAccessToken = errors.New("content: github releases: missing access token")
	ErrMissingRepoOwner   = errors.New("content: github releases: missing repo owner")
	ErrMissingRepoName    = errors.New("content: github releases: missing repo name")
	ErrMissingRecipient   = errors.New("content: github releases: missing recipient")
)

type Error struct {
	Inner []error `json:"errors"`
}

func (e *Error) Error() string {
	errStrings := make([]string, len(e.Inner))
	for i, err := range e.Inner {
		errStrings[i] = err.Error()
	}
	return strings.Join(errStrings, "; ")
}
