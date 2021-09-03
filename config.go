package main

import (
	"github.com/simp7/patent-middle-server/storage/cache"
	"github.com/simp7/patent-middle-server/storage/rest"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Rest    rest.Config  `yaml:"rest-server"`
	Cache   cache.Config `yaml:"database"`
	Port    int          `yaml:"port"`
	Version string       `yaml:"version"`
}

func getSkelConfig() (Config, error) {
	return getConfigFrom(skelTo("config.yaml"))
}

func GetConfig() (Config, error) {
	return getConfigFrom(rootTo("config.yaml"))
}

func SetConfig(config Config) (err error) {

	var data []byte

	if data, err = yaml.Marshal(config); err == nil {
		err = os.WriteFile(rootTo("config.yaml"), data, 0644)
	}

	return

}

func getConfigFrom(path string) (conf Config, err error) {

	var file *os.File

	if file, err = os.Open(path); err == nil {
		defer file.Close()
		err = yaml.NewDecoder(file).Decode(&conf)
	}

	return

}

func IsLatest() bool {

	realConf, _ := GetConfig()
	skelConf, _ := getSkelConfig()

	return realConf.Version == skelConf.Version

}

func UpdateVersion() error {

	realConf, _ := GetConfig()
	skelConf, _ := getSkelConfig()

	realConf.Version = skelConf.Version
	return SetConfig(realConf)

}
