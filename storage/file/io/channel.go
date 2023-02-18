package io

type Channel interface {
	Read(buf Buffer, position int64) error

	Write(buf Buffer, position int64) error

	Append(buf Buffer) (int64, error)

	Size() (int64, error)

	Close() error
}
