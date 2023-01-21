package sql

// Record a map from field names to constants
type Record interface {
	GetVal(fldName string) Constant
}
