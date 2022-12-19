package sql

type Type interface {

	// GetSqlType Returns the SQL type corresponding to this instance.
	GetSqlType() int

	// GetArgument Returns the argument associated with this instance.
	// For example, this methods returns 20 for the SQL type VARCHAR(20).
	GetArgument() int

	// IsFixedSize Returns whether the number of bytes required to encode
	IsFixedSize() bool

	// IsNumeric Returns whether the values of this type is numeric.
	IsNumeric() bool

	// MaxSize Returns the maximum number of bytes required to encode
	MaxSize() int

	MaxValue()

	MinValue()
}

// NewInstance Constructs a new instance corresponding
// to the specified SQL type and argument.
func NewInstance(sqlType, arg int) Type {

	// Unsupported SQL type
	return nil
}
