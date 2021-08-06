package main

import (
	"github.com/simp7/patent-middle-server/storage"
	"github.com/simp7/patent-middle-server/storage/cache"
	"github.com/simp7/patent-middle-server/storage/rest"
	"log"
	"os"
)

func main() {

	cacheDB, err := cache.Mongo("mongodb://localhost")
	if err != nil {
		log.Println("Can't connect Database. Change server to No-Cache-mode.")
		cacheDB = cache.Nocache()
	}
	source := rest.New("http://plus.kipris.or.kr/kipo-api/kipi/patUtiModInfoSearchSevice/getWordSearch", "http://plus.kipris.or.kr/kipo-api/kipi/patUtiModInfoSearchSevice/getBibliographyDetailInfoSearch", os.Getenv("KIPRIS"))

	s := New(80, storage.New(source, cacheDB))
	defer s.Close()

	err = s.Start()
	log.Fatal(err)

}
