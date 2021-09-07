package main

import (
	"github.com/simp7/patent-middle-server/config"
	"github.com/simp7/patent-middle-server/model"
	"os"
	"os/exec"
)

type FileSystem interface {
	OpenLogfile() (*os.File, error)
	SaveCSVFile(group *model.CSVGroup) (file *os.File, err error)
	RemoveCSVFile(group *model.CSVGroup) error
	LoadConfig() (config config.Config, err error)
	Word2vec(args ...string) *exec.Cmd
	LDA(args ...string) *exec.Cmd
	LSA(args ...string) *exec.Cmd
}
