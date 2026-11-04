package ksqlite

import "fmt"

var _ error = &ErrUnknownToken{}

type ErrUnknownToken struct {
	Content string
}

func (e ErrUnknownToken) Error() string {
	return fmt.Sprintf("unknown token %s", e.Content)
}

