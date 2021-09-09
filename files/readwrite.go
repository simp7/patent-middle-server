package files

import (
	"io"
	"os"
	"os/exec"
)

type ReadWrite interface {
	Open(file Path, isAppend bool) (*os.File, error)
	Mkdir(dir Path) error
	Write(file Path, data []byte) error
	IsExist(file Path) bool
	Command(file Path, args ...string) *exec.Cmd
	GetUpdateList() (result []string, err error)
	Remove(file Path) error
	WriteWithChan(file Path, stream <-chan string) (err error)
	Copy(to Path, from io.Reader) (err error)
}
