package files

import (
	"io/fs"
)

type ReadOnly interface {
	ReadDir(dir Path) ([]fs.DirEntry, error)
	Open(file Path) (fs.File, error)
}
