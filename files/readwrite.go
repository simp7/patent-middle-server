package files

import (
	"os"
	"os/exec"
)

type ReadWrite interface {
	Open(file Path) (*os.File, error)
	Create(file Path) (*os.File, error)
	Mkdir(dir Path) error
	Write(file Path, data []byte) error
	IsExist(file Path) bool
	Command(file Path, args ...string) *exec.Cmd
	GetUpdateList() (result []string, err error)
	Remove(file Path) error
}
