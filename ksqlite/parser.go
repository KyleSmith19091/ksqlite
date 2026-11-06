package ksqlite

import (
	"fmt"

	"github.com/xwb1989/sqlparser"
)

type Parser struct {

}

func NewParser() *Parser {
	return &Parser{}
}

func Parse(sql string) error {
	statement, err := sqlparser.Parse(sql)
	if err != nil {
		return fmt.Errorf("could not parse SQL: %w", err)
	}

	switch statement := statement.(type) {
		case *sqlparser.Select:
			statement.SelectExprs				
	}
}
