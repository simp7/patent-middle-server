package main

import (
	"github.com/simp7/patent-middle-server/config"
	"github.com/simp7/patent-middle-server/model"
	"io"
)

type FileSystem interface {
	BindLogFiles(path ...string) (io.WriteCloser, error)
	SaveCSVFile(group *model.CSVGroup) error
	RemoveCSVFile(group *model.CSVGroup) error
	LoadConfig() (config config.Config, err error)
	Word2vec(args ...string) ([]byte, error)
	LDA(args ...string) ([]byte, error)
	LSA(args ...string) ([]byte, error)
}
