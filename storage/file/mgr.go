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
