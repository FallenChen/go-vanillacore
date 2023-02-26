package file

import (
	"errors"
	fileIO "github.com/go-vanillacore/storage/file/io"
	"github.com/go-vanillacore/storage/log"
	util "github.com/go-vanillacore/util/sync"
	"hash/fnv"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

const (
	Db_Files_Dir         string = "db_files"
	Log_Files_Dir        string = "log_files"
	Tmp_File_Name_Prefix        = "_temp"
	length               int32  = 1009
)

type Mgr struct {
	DbDirectory, LogDirectory *os.File

	IsNew bool

	OpenFiles util.Map[string, *fileIO.Channel]

	FileNotEmptyCache util.Map[string, bool]

	anchors [length]string

	lock [length]*sync.Mutex
}

func NewMgr(dbName string) (*Mgr, error) {
	dbFile := filepath.Join(Db_Files_Dir, dbName)
	dbDirectory, err := os.Create(dbFile)
	if err != nil {
		return nil, err
	}

	logFile := filepath.Join(Log_Files_Dir, dbName)
	logDirectory, err := os.Create(logFile)
	if err != nil {
		return nil, err
	}

	isNew := !fileExists(dbFile)

	// check the existence of log folder
	if !isNew && !fileExists(logFile) {
		return nil, errors.New("log file for the existed " + dbName + " is missing")
	}

	// create the directory if the database is new
	if isNew && (!mkdir(dbFile)) {
		return nil, errors.New("cannot create " + dbName)
	}

	// remove any leftover temporary tables
	dirs, err := os.ReadDir(dbFile)
	if err != nil {
		return nil, err
	}
	for _, file := range dirs {
		if strings.HasPrefix(file.Name(), Tmp_File_Name_Prefix) {
			join := filepath.Join(dbFile, file.Name())
			err = os.Remove(join)
			if err != nil {
				return nil, err
			}
		}
	}
	mgr := &Mgr{
		DbDirectory:  dbDirectory,
		LogDirectory: logDirectory,
		IsNew:        isNew,
	}
	return mgr, nil
}

func mkdir(dbName string) bool {
	err := os.Mkdir(dbName, 0755)
	return err == nil
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func (file *Mgr) prepareAnchor(anchor string) *sync.Mutex {
	h := fnv.New32()
	h.Write([]byte(anchor))
	hashCode := int32(h.Sum32())

	code := hashCode % length

	if code < 0 {
		code += length
	}
	return file.lock[code]
}

// Reads the contents of a disk block into a byte buffer.
func (file *Mgr) read(blk BlockID, buffer fileIO.Buffer) error {
	fileChannel, err := file.getFileChannel(blk.FileName)

	if err != nil {
		return err
	}
	// clear the buffer
	buffer.Clear()

	// read a block from file
	return fileChannel.Read(buffer, blk.BlkNum*BlockSize)

	// for controller
}

func (file *Mgr) write(blk BlockID, buffer fileIO.Buffer) error {
	fileChannel, err := file.getFileChannel(blk.FileName)

	if err != nil {
		return err
	}

	// clear the buffer
	buffer.Rewind()

	// read a block from file
	return fileChannel.Write(buffer, blk.BlkNum*BlockSize)

	// for controller
}

func (file *Mgr) append(fileName string, buffer fileIO.Buffer) (*BlockID, error) {
	fileChannel, err := file.getFileChannel(fileName)
	if err != nil {
		return nil, err
	}

	blkNum, err := fileChannel.Append(buffer)
	if err != nil {
		return nil, err
	}

	return New(fileName, blkNum), nil
}

func (file *Mgr) Size(fileName string) (int64, error) {
	channel, err := file.getFileChannel(fileName)
	if err != nil {
		errors.New("cannot access " + fileName)
	}
	size, err := channel.Size()
	return size / BlockSize, err
}

func (file *Mgr) getFileChannel(fileName string) (*fileIO.Channel, error) {
	mutex := file.prepareAnchor(fileName)
	mutex.Lock()
	defer mutex.Unlock()
	fileChannel, ok := file.OpenFiles.Load(fileName)

	if !ok {
		if fileName == log.Default_Log_File {
			dbFile, err := createChildFile(file.LogDirectory, fileName)
			if err != nil {
				return nil, err
			}
			channel, err := fileIO.NewIoChannel(dbFile)
			if err != nil {
				return nil, err
			}
			fileChannel = fileIO.NewVirtualChannel(channel)
			//file.OpenFiles.Store(fileName, fileChannel)
		}
	}

	return fileChannel, nil
}

func createChildFile(parent *os.File, fileName string) (*os.File, error) {
	if fileName == "" {
		return nil, errors.New("fileName is empty")
	}

	path := filepath.Join(parent.Name(), fileName)

	// 如果文件已经存在，则返回已有的文件对象
	if _, err := os.Stat(path); err == nil {
		f, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		return f, nil
	}

	// 创建新的文件
	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}

	return f, nil
}
