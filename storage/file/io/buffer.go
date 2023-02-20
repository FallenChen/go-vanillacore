package io

type Buffer interface {
	Get(position int, dst []byte) Buffer

	Put(position int, src []byte) Buffer

	Clear()

	Rewind()

	Close()
}
