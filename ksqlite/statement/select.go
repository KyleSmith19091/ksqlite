package statement

import "github.com/KyleSmith19091/sqlite/storage"

var _ Statement = &Select{}

type Select struct {
	Expr []SelectExpression
	From string
	Where string
}

func (s *Select) Execute(engine storage.Engine) error {
	// translate from SQL to storage engine instructions
	
	// query planner to translate table, column names and conditions to storage engine calls

	// Read(table, columns, conditions)

	return nil
}

type SelectExpression interface {
}

type ColumnName struct {
	Name string
}

type Star struct {
}


