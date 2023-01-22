package sql

import (
	"database/sql"
	"database/sql/driver"
	"github.com/go-vanillacore/util"
)

type Int32Type struct {
}

func NewInt32Type() *Int32Type {
	return &Int32Type{}
}

func (i Int32Type) GetSqlType() driver.Valuer {
	return sql.NullInt32{}
}

func (i Int32Type) GetArgument() int8 {
	return -1
}

func (i Int32Type) IsFixedSize() bool {
	return true
}

func (i Int32Type) IsNumeric() bool {
	return true
}

func (i Int32Type) MaxSize() int8 {
	return util.INT32SIZE
}

func (i Int32Type) MaxValue() Constant {
	return nil
}

func (i Int32Type) MinValue() Constant {
	return nil
}
