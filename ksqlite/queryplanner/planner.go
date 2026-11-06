package queryplanner

import (
	"fmt"
	"reflect"
	"slices"
)

type Planner interface {
	PlanSelect(SelectPlanInput) error
}

type SelectPlanInput struct {
	TableName string
	ColumnNames []string
}

var _ Planner = &PlannerImpl{}

type Column struct {
	Name string
	Type reflect.Type
}

type Table struct {
	Name string
	Columns []Column
}

type Row struct {
	// contiguous array of bytes
	buffer []byte
	
	// each element represents an offset to index
	// into the buffer
	columnOffsets []struct{
		offset uint32
		ttype reflect.Type
	}
}

type PlannerImpl struct {
	tables map[string]Table
}

func (p *PlannerImpl) PlanSelect(sInput SelectPlanInput) error {
	table, found := p.tables[sInput.TableName]
	if !found {
		return fmt.Errorf("table not found")
	}

	numColumns := len(sInput.ColumnNames)
	if sInput.ColumnNames[0] == "*" {
		numColumns = len(table.Columns)
	}

	// recording columns to be read
	columns := make([]Column, len(sInput.ColumnNames))
	for nameIdx, column := range sInput.ColumnNames {
		// find index of column in table
		colIdx := slices.IndexFunc(table.Columns, func(col Column) bool { return col.Name == column }) 
		if colIdx == -1 {
			return fmt.Errorf("column not found for table")
		}

		columns[nameIdx] = Column{
			Name: column,
			Type: table.Columns[colIdx].Type,
		}
	}

	// Ok I know which columns to read, now what?
	// Ideally want a plan that can be handed to an engine?
	// Plan says I want to read these columns, from this table, and filter on this condition
	

	// What is a table? How is it laid out

	
	return nil
}

