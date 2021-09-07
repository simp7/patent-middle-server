package subsystem

import (
	"github.com/simp7/patent-middle-server/files"
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

func (r *realSystem) Open(file files.Path) (*os.File, error) {
	return os.OpenFile(r.path(file), os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
}

func (r *realSystem) Create(file files.Path) (created *os.File, err error) {
	if created, err = os.Create(r.path(file)); err == nil {
		err = os.Chmod(r.path(file), 0755)
	}
	return
}

func (r *realSystem) Mkdir(dir files.Path) error {
	return os.Mkdir(r.path(dir), 0700)
}

func (r *realSystem) Write(file files.Path, data []byte) error {
	return os.WriteFile(r.path(file), data, 0644)
}

func (r *realSystem) IsExist(file files.Path) bool {
	_, err := os.Stat(r.path(file))
	return !os.IsNotExist(err)
}

func (r *realSystem) Command(file files.Path, args ...string) *exec.Cmd {

	if file == files.WORD2VEC || file == files.LDA || file == files.LSA {
		totalArgs := []string{r.path(file)}
		return exec.Command(r.path(files.EXECUTE), append(totalArgs, args...)...)
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

func (r *realSystem) path(file files.Path) string {
	return path.Join(os.Getenv("HOME"), "patent-server", string(file))
}
