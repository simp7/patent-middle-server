package main

import (
	"github.com/simp7/patent-middle-server/storage/cache"
	"github.com/simp7/patent-middle-server/storage/rest"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Rest  rest.Config  `yaml:"rest-server"`
	Cache cache.Config `yaml:"database"`
	Port  int          `yaml:"port"`
}

func getConfig() (conf Config, err error) {

	var file *os.File

	file, err = os.Open(rootTo("config.yaml"))
	if err != nil {
		return
	}

	err = yaml.NewDecoder(file).Decode(&conf)

	return

}
