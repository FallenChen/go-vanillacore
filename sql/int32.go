package sql

import (
	"database/sql"
	"database/sql/driver"
)

type Int32Type struct {
}

func (i Int32Type) GetSqlType() driver.Valuer {
	return sql.NullInt32{}
}

func (i Int32Type) GetArguments() int8 {
	return -1
}

func (i Int32Type) IsFixedSize() bool {
	return true
}
