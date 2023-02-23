package file

import (
	"github.com/go-vanillacore/sql"
	"github.com/go-vanillacore/storage/file/io"
	"github.com/go-vanillacore/util"
)

const BlockSize int64 = 4096

type Page struct {
	contents io.Buffer
	fileMgr  *Mgr
}

func MaxSize(t sql.Type) int32 {
	if t.IsFixedSize() {
		return t.MaxSize()
	} else {
		return util.Int32Size + t.MaxSize()
	}
}

func Size(val sql.Constant) int32 {
	if val.GetType().IsFixedSize() {
		return val.Size()
	} else {
		return util.Int32Size + val.Size()
	}
}

func (page *Page) Close() {
	page.contents.Close()
}
