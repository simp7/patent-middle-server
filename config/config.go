package config

import (
	"github.com/simp7/patent-middle-server/storage/cache"
	"github.com/simp7/patent-middle-server/storage/rest"
)

type Config struct {
	Rest    rest.Config  `yaml:"rest-server"`
	Cache   cache.Config `yaml:"database"`
	Port    int          `yaml:"port"`
	Version string       `yaml:"version"`
}
