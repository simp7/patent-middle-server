package main

import (
	"github.com/google/logger"
	"github.com/simp7/patent-middle-server/files"
	"github.com/simp7/patent-middle-server/files/subsystem"
	"github.com/simp7/patent-middle-server/storage"
	"github.com/simp7/patent-middle-server/storage/cache"
	"github.com/simp7/patent-middle-server/storage/rest"
	"log"
)

func main() {

	fs, err := files.System(subsystem.Real(), subsystem.Skel())
	if err != nil {
		log.Fatal(err)
	}

	conf, err := fs.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	logWriter, _ := fs.BindLogFiles()
	serverLogger := logger.Init("server", true, false, logWriter)
	restLogger := logger.Init("rest", true, false, logWriter)
	cacheLogger := logger.Init("cache", true, false, logWriter)

	cacheDB := newCacheDB(conf.Cache, cacheLogger)
	source := rest.New(conf.Rest, restLogger)

	middleServer := New(conf.Port, storage.New(source, cacheDB), fs, serverLogger)
	err = middleServer.Start()
	defer middleServer.Close()

}

func newCacheDB(conf cache.Config, cacheLogger *logger.Logger) storage.Cache {

	cacheDB, err := cache.Mongo(conf, cacheLogger)
	if err != nil {
		cacheLogger.Error(err)
		cacheLogger.Info("change server to No-Cache-mode")
		cacheDB = cache.Nocache()
	}

	return cacheDB

}
