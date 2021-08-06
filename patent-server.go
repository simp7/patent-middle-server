package main

import (
	"github.com/simp7/patent-middle-server/claimStorage"
	"github.com/simp7/patent-middle-server/claimStorage/cache"
	"log"
	"os"
)

func main() {

	cacheDB, err := cache.Mongo("mongodb://localhost")
	if err != nil {
		log.Println("Can't connect Database. Change server to No-Cache-mode.")
		cacheDB = cache.Nocache()
	}
	storage := claimStorage.New("http://plus.kipris.or.kr/kipo-api/kipi/patUtiModInfoSearchSevice/getWordSearch", "http://plus.kipris.or.kr/kipo-api/kipi/patUtiModInfoSearchSevice/getBibliographyDetailInfoSearch", os.Getenv("KIPRIS"), cacheDB)

	s := New(80, storage)
	defer s.Close()

	err = s.Start()
	log.Fatal(err)

}
