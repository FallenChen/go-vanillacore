package util

import (
	"bytes"
	"encoding/binary"
)

const (
	Int32Size = int32(32 / 8)
)

func ToBytes(val int32) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, val)
	return bytesBuffer.Bytes()
}
