package main

import (
	"embed"
	_ "embed"
	"github.com/simp7/patent-middle-server/storage"
	"github.com/simp7/patent-middle-server/storage/cache"
	"github.com/simp7/patent-middle-server/storage/rest"
	"gopkg.in/yaml.v3"
	"io"
	"io/fs"
	"log"
	"os"
)

//go:embed skel
var skel embed.FS

func main() {

	conf, err := getConfig()
	if err != nil && err != io.EOF {
		panic(err)
	}

	cacheDB := newCacheDB(conf.Cache)
	source := rest.New(conf.Rest)

	middleServer := New(conf.Port, storage.New(source, cacheDB))
	defer middleServer.Close()

	err = middleServer.Start()
	log.Fatal(err)

}

func getConfig() (conf Config, err error) {

	home := os.Getenv("HOME")
	file, err := os.Open(home + "/patent-server/config.yaml")

	if err != nil {

		var skelFile fs.File
		var created *os.File

		err = os.Mkdir(home+"/patent-server", 0700)
		if err != nil {
			return
		}

		skelFile, err = skel.Open("skel/config.yaml")
		if err != nil {
			return
		}

		created, err = os.Create(home + "/patent-server/config.yaml")
		if err != nil {
			return
		}

		_, err = io.Copy(created, skelFile)
		if err != nil && err != io.EOF {
			return
		}

		file, err = os.Open(home + "/patent-server/config.yaml")

	}

	err = yaml.NewDecoder(file).Decode(&conf)
	return

}

func newCacheDB(conf cache.Config) storage.Cache {

	cacheDB, err := cache.Mongo(conf)
	if err != nil {
		log.Println("Can't connect Database. Change server to No-Cache-mode.")
		cacheDB = cache.Nocache()
	}

	return cacheDB

}
