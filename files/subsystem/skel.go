package subsystem

import (
	"embed"
	"github.com/simp7/patent-middle-server/files"
	"io/fs"
	"path"
)

//go:embed skel
var fileSystem embed.FS

type skelSystem struct {
}

func Skel() files.ReadOnly {
	return &skelSystem{}
}

func (s *skelSystem) ReadDir(path files.Path) ([]fs.DirEntry, error) {
	return fileSystem.ReadDir(s.path(path))
}

func (s *skelSystem) Open(path files.Path) (fs.File, error) {
	file, err := fileSystem.Open(s.path(path))
	return file, err
}

func (s *skelSystem) path(file files.Path) string {
	return path.Join("skel", string(file))
}
