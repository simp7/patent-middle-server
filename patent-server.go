package main

import (
	"github.com/simp7/patent-middle-server/storage"
	"github.com/simp7/patent-middle-server/storage/cache"
	"github.com/simp7/patent-middle-server/storage/rest"
	"log"
)

func main() {

	err := initialize()
	if err != nil {
		log.Fatal(err)
	}

	conf, err := getConfig()
	if err != nil {
		log.Fatal(err)
	}

	cacheDB := newCacheDB(conf.Cache)
	source := rest.New(conf.Rest)

	middleServer := New(conf.Port, storage.New(source, cacheDB))
	defer middleServer.Close()

	err = middleServer.Start()
	log.Fatal(err)

}

func newCacheDB(conf cache.Config) storage.Cache {

	cacheDB, err := cache.Mongo(conf)
	if err != nil {
		log.Println("Can't connect Database. Change server to No-Cache-mode.")
		cacheDB = cache.Nocache()
	}

	return cacheDB

}
