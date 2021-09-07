package main

import (
	"github.com/google/logger"
	"github.com/simp7/patent-middle-server/files"
	"github.com/simp7/patent-middle-server/files/subsystem"
	"github.com/simp7/patent-middle-server/logWriter"
	"github.com/simp7/patent-middle-server/storage"
	"github.com/simp7/patent-middle-server/storage/cache"
	"github.com/simp7/patent-middle-server/storage/rest"
	"log"
)

func main() {

	sys, err := files.System(subsystem.Real(), subsystem.Skel())
	if err != nil {
		log.Fatal(err)
	}

	conf, err := sys.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	cacheDB := newCacheDB(conf.Cache)
	source := rest.New(conf.Rest)

	logFile, err := sys.OpenLogfile()
	if err != nil {
		logFile = nil
	}

	lg := logger.Init("server", true, false, logWriter.New(logFile))

	middleServer := New(conf.Port, storage.New(source, cacheDB), lg, sys)
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
