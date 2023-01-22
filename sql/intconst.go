package sql

type Int32Constant struct {
	val int32
}

func GetType() Type {
	return NewInt32Type()
}
