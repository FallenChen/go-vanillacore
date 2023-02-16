package util

import (
	"bytes"
	"encoding/binary"
)

const (
	INT32SIZE = 4
)

func ToBytes(val int32) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, val)
	return bytesBuffer.Bytes()
}
