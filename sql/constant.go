package sql

// Constant denotes a value of a supported type
type Constant interface {

	// GetType Returns the type corresponding to this constant.
	GetType() Constant

	// AsGoVal Returns the Go object corresponding to this constant.
	AsGoVal()

	AsBytes()

	Size() int

	CastTo(p Type) Constant

	Add(c Constant) Constant

	Sub(c Constant) Constant

	Mul(c Constant) Constant

	Div(c Constant) Constant
}

func DefaultInstance(t Type) Constant {
	v := t.GetSqlType()
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
	return nil
}
