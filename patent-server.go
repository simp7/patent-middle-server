package main

import (
	"github.com/google/logger"
	"github.com/simp7/patent-middle-server/logWriter"
	"github.com/simp7/patent-middle-server/storage"
	"github.com/simp7/patent-middle-server/storage/cache"
	"github.com/simp7/patent-middle-server/storage/rest"
	"log"
	"os"
)

func main() {

	if err := initialize(); err != nil {
		log.Fatal(err)
	}

	conf, err := GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	cacheDB := newCacheDB(conf.Cache)
	source := rest.New(conf.Rest)

	logFile, err := os.OpenFile(rootTo("server.log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		logFile = nil
	}

	lg := logger.Init("server", true, false, logWriter.New(logFile))

	middleServer := New(conf.Port, storage.New(source, cacheDB), lg)
	defer middleServer.Close()

	err = middleServer.Start()
	log.Fatal(err)

}

func newCacheDB(conf cache.Config) storage.Cache {

	cacheDB, err := cache.Mongo(conf)
	if err != nil {
		log.Println(err)
		log.Println("Change server to No-Cache-mode.")
		cacheDB = cache.Nocache()
	}

	return cacheDB

}
