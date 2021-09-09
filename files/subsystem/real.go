package subsystem

import (
	"github.com/simp7/patent-middle-server/files"
	"io"
	"os"
	"os/exec"
	"path"
	"strings"
)

type realSystem struct {
}

func Real() files.ReadWrite {
	return &realSystem{}
}

func (r *realSystem) Open(file files.Path, isAppend bool) (*os.File, error) {

	mode := os.O_RDWR | os.O_CREATE
	if isAppend {
		mode = mode | os.O_APPEND
	}

	return os.OpenFile(r.path(file), mode, 0755)

}

func (r *realSystem) Mkdir(dir files.Path) error {
	return os.Mkdir(r.path(dir), 0700)
}

func (r *realSystem) Write(file files.Path, data []byte) error {

	opened, err := r.Open(file, false)
	if err != nil {
		return err
	}
	defer opened.Close()

	_, err = opened.Write(data)
	return err

}

func (r *realSystem) IsExist(file files.Path) bool {
	_, err := os.Stat(r.path(file))
	return !os.IsNotExist(err)
}

func (r *realSystem) Command(file files.Path, args ...string) *exec.Cmd {

	if file == files.WORD2VEC || file == files.LDA || file == files.LSA {
		return r.nlp(file, args...)
	}

	return exec.Command(r.path(file), args...)

}

func (r *realSystem) GetUpdateList() (result []string, err error) {

	var data []byte

	findCmd := exec.Command("find", r.path(files.ROOT), "-type", "directory", "-name", "venv", "-prune", "-and", "!", "-name", "venv", "-o", "-type", "file", "-and", "!", "-name", "*.log", "-and", "!", "-name", "config.*")
	if data, err = findCmd.Output(); err == nil {
		result = strings.Split(string(data), "\n")
	}

	return

}

func (r *realSystem) Remove(file files.Path) error {
	return os.Remove(r.path(file))
}

func (r *realSystem) WriteWithChan(file files.Path, stream <-chan string) (err error) {

	var opened *os.File

	if opened, err = r.Open(file, true); err == nil {
		for line := range stream {
			_, err = opened.WriteString(line)
		}
		err = opened.Close()
	}

	return

}

func (r *realSystem) Copy(file files.Path, source io.Reader) (err error) {

	var created *os.File
	if created, err = r.Open(file, false); err != nil {
		return
	}

	if _, err = io.Copy(created, source); err != nil && err != io.EOF {
		return
	}

	err = created.Close()
	return

}

func (r *realSystem) path(file files.Path) string {
	return path.Join(os.Getenv("HOME"), "patent-server", string(file))
}

func (r *realSystem) nlp(file files.Path, arg ...string) *exec.Cmd {

	copied := make([]string, len(arg))
	copy(copied, arg)
	copied[0] = r.path(files.New(arg[0]))

	result := []string{r.path(file)}
	result = append(result, copied...)

	return exec.Command(r.path(files.EXECUTE), result...)

}
