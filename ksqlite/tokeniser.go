package ksqlite

import (
	"strings"

	"golang.org/x/text/runes"
)

type TokenID uint8

type TokenType struct {
	Id        TokenID
	IsKeyword bool
	Content   string
}

const (
	TOKEN_ID_SELECT = iota
	TOKEN_ID_INSERT
	TOKEN_ID_INTO
	TOKEN_ID_LITERAL
)

var registeredTokenIndex = []TokenType{
	{
		Id:        TOKEN_ID_SELECT,
		IsKeyword: true,
		Content:   "select",
	},
	{
		Id:        TOKEN_ID_INSERT,
		IsKeyword: true,
		Content:   "insert",
	},
	{
		Id:        TOKEN_ID_INTO,
		IsKeyword: true,
		Content:   "into",
	},
}
var TOKENS = map[string]TokenType{}

func init() {
	for _, tokenType := range registeredTokenIndex {
		TOKENS[tokenType.Content] = tokenType
	}
}

type Token struct {
	Type      TokenID
	IsKeyword bool
	Content   string
}

// number literal string
func Tokenise(input string) ([]Token, error) {
	tokens := []Token{}

	if len(input) == 0 {
		return []Token{}, nil
	}	

	char := rune(input[0])
	if isNumber(char) {
		handleNumber()
	}
	if isLetter(char) {
		// handle keyword
	}

}

func handleNumber(input *string) (Token, error) {

}

func consume(num uint32, input* string) error {
	// check if we are trying to consume more characters than we have string left
	if len(*input) < int(num) {
	}
}

func isLetter(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

func isNumber(r rune) bool {
	return r >= '0' && r <= '9'
}
