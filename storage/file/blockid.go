package file

import (
	"hash/fnv"
	"strconv"
	"strings"
)

type BlockID struct {
	FileName   string
	BlkNum     uint64
	MyHashCode uint32
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

func New(fileName string, blkNum uint64) BlockID {
	bolckID := BlockID{FileName: fileName, BlkNum: blkNum}
	h := fnv.New32()
	h.Write([]byte(bolckID.ToString()))
	bolckID.MyHashCode = h.Sum32()
	return bolckID
}

func (b BlockID) ToString() string {
	return "[file " + b.FileName + ", block " + strconv.FormatUint(b.BlkNum, 10) + "]"
}

func (b BlockID) HashCode() uint32 {
	return b.MyHashCode
}
