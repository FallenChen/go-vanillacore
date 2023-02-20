package file

import (
	"github.com/go-vanillacore/storage/file/io"
	"github.com/go-vanillacore/util/sync"
	"hash/fnv"
	"os"
)

const (
	Db_Files_Dir         string = "db_files"
	Log_Files_Dir        string = "log_files"
	Tmp_File_Name_Prefix        = "_temp"
	length               int32  = 1009
)

type FileMgr struct {
	dbDirectory, logDirectory os.File

	isNew bool

	openFiles sync.Map[string, io.Channel]

	fileNotEmptyCache sync.Map[string, bool]

	anchors [length]string
}

func (file *FileMgr) prepareAnchor(anchor string) string {
	h := fnv.New32()
	h.Write([]byte(anchor))
	hashCode := int32(h.Sum32())

	code := hashCode % length

	if code < 0 {
		code += length
	}
	return file.anchors[code]
}

func read(blk BlockID, buffer io.Buffer) {
	fileChannel := getFileChannel(blk.FileName)

	// clear the buffer
	buffer.Clear()

	// read a block from file
	fileChannel.Read(buffer, blk.BlkNum*BlockSize)

	// for controller
}

func write(blk BlockID, buffer io.Buffer) {
	fileChannel := getFileChannel(blk.FileName)

	// clear the buffer
	buffer.Rewind()

	// read a block from file
	fileChannel.Write(buffer, blk.BlkNum*BlockSize)

	// for controller
}

func append(fileName string, buffer io.Buffer) (*BlockID, error) {
	fileChannel := getFileChannel(fileName)

	blkNum, err := fileChannel.Append(buffer)
	if err != nil {
		return nil, err
	}

	return New(fileName, blkNum), nil
}

func getFileChannel(fileName string) io.Channel {
	return nil
}
