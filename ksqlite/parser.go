package ksqlite

import "errors"

type ASTNode struct {
}

func Parse(tokens []Token) error {
	if len(tokens) == 0 {
		return nil
	}

	switch tokens[0].Type {
	case TOKEN_ID_SELECT:
		return handleSelect(tokens[1:])
	case TOKEN_ID_INSERT:
		return handleInsert(tokens[1:])
	default:
		return errors.New("expected a top-level keyword")
	}
}

func handleSelect(args []Token) error {
	if len(args) == 0 {
		return errors.New("select statement requires arguments, but none were provided")
	}

	return nil
}

func handleInsert(args []Token) error {
	if len(args) == 0 {
		return errors.New("select statement requires arguments, but none were provided")
	}

	// insert into table_name values 1 2 "hello world"

	return nil
}
