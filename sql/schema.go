package sql

type Schema struct {
	Fields map[string]*Type
}

// addField adds a field to this schema having a specified name and type
func (schema *Schema) addField(name string, fieldType *Type) {

}

// add adds a field in another schema having the specified name to this schema
func (schema *Schema) add(fldName string, sch *Schema) {

}

// addAll adds of the fields in the specified schema to this schema
func (schema *Schema) addAll(sch *Schema) {

}

// fields returns a sorted set containing the field names in this schema, sorted by
// their natural ordering
func (schema *Schema) fields() {

}
