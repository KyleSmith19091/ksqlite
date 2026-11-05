package ksqlite

import (
	"fmt"
)

type TokenID uint8

type TokenType struct {
	Id        TokenID
	IsKeyword bool
	Content   string
}

const (
	TOKEN_ID_SELECT TokenID = iota
	TOKEN_ID_INSERT
	TOKEN_ID_INTO
	TOKEN_ID_LITERAL
	TOKEN_ID_INTEGER
	TOKEN_ID_FLOAT
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

var characterLookupTable = map[rune]rune{}

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

	for len(input) != 0 {
		char := rune(input[0])
		if char == '-' || isDigit(char) {
			token, err := handleNumber(&input)
			if err != nil {
				return nil, fmt.Errorf("could not tokenise number: %w", err)
			}
			tokens = append(tokens, token)	
		} else if isLetter(char) {
		}
	}

	return nil, nil
}

func consume(num uint32, input *string) (string, bool, error) {
	// check if we are trying to consume more characters than we have string left
	if len(*input) < int(num) {
		return "", false, fmt.Errorf("consuming exceeds input buffer")
	}

	inputPrefix := (*input)[:num]
	*input = (*input)[num:]

	return inputPrefix, len(*input) == 0, nil
}

func peek(input *string) rune {
	return rune((*input)[0])
}

func handleLetter(input *string) (Token, error) {
	letter, empty, err := consume(1, input)
	if err != nil {
		return Token{}, fmt.Errorf("could not consume first letter: %w", err)
	}

	if empty {
		return Token{
			Type: TOKEN_ID_LITERAL,
			Content: letter,
		}, nil
	}

	for char := peek(input); isLetter(char); {
	}

	return Token{}, nil
}

func handleNumber(input *string) (Token, error) {
	number, empty, err := consume(1, input)
	if err != nil {
		return Token{}, fmt.Errorf("could not consume first digit: %w", err)
	}

	if empty && number == "-" {
		return Token{}, fmt.Errorf("invalid number, can't just contain -")
	}

	if empty {
		return Token{
			Type: TOKEN_ID_INTEGER,
			Content: number,
		}, nil
	}

	decimalPointFound := false
	for char := peek(input); len(*input) != 0 && (char == '.' || isDigit(char)); {
		if char == '.' && decimalPointFound {
			return Token{}, fmt.Errorf("found more than one . in number token")
		}

		// found decimal point, set flag
		if char == '.' {
			decimalPointFound = true
		}

		// consume character from input
		nextTokenPart, _, err := consume(1, input)
		if err != nil {
			return Token{}, fmt.Errorf("could not consume part of number: %w", err)
		}

		// append to current number token
		number += nextTokenPart
	}

	tokenID := TOKEN_ID_INTEGER
	if decimalPointFound {
		tokenID = TOKEN_ID_FLOAT
	}
	
	return Token{
		Type: tokenID,
		Content: number,
	}, nil
}

func isLetter(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}
