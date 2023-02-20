package file

import (
	"hash/fnv"
	"strconv"
	"strings"
)

type BlockID struct {
	FileName   string
	BlkNum     int64
	MyHashCode int32
}

func (b BlockID) CompareTo(other BlockID) int {
	nameResult := strings.Compare(b.FileName, other.FileName)
	if nameResult != 0 {
		return nameResult
	}

	if b.BlkNum < other.BlkNum {
		return -1
	} else if b.BlkNum > other.BlkNum {
		return 1
	}
	return 0
}

func New(fileName string, blkNum int64) *BlockID {
	bolckID := &BlockID{FileName: fileName, BlkNum: blkNum}
	h := fnv.New32()
	h.Write([]byte(bolckID.ToString()))
	bolckID.MyHashCode = int32(h.Sum32())
	return bolckID
}

func (b BlockID) ToString() string {
	return "[file " + b.FileName + ", block " + strconv.FormatInt(b.BlkNum, 10) + "]"
}

func (b BlockID) HashCode() int32 {
	return b.MyHashCode
}
