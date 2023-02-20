package sql

import (
	"database/sql/driver"
)

type Type interface {

	// GetSqlType Returns the SQL type corresponding to this instance.
	GetSqlType() driver.Valuer

	// GetArgument Returns the argument associated with this instance.
	// For example, these methods returns 20 for the SQL type VARCHAR(20).
	GetArgument() int8

	// IsFixedSize Returns whether the number of bytes required to encode
	IsFixedSize() bool

	// IsNumeric Returns whether the values of this type is numeric.
	IsNumeric() bool

	// MaxSize Returns the maximum number of bytes required to encode
	MaxSize() int32

	MaxValue() Constant

	MinValue() Constant
}

// NewInstance Constructs a new instance corresponding
// to the specified SQL type and argument.
func NewInstance(v driver.Valuer, arg int) Type {
	switch v.(type) {
	//case sql.NullInt32:
	//	return
	//case sql.NullInt64:
	//	return
	//case sql.NullFloat64:
	//	return
	//case sql.NullString:
	//	return
	}
	// Unsupported SQL type
	return nil
}
