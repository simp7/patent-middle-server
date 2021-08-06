package main

import (
	"fmt"
	"github.com/simp7/patent-middle-server/claimStorage"
	"github.com/simp7/patent-middle-server/claimStorage/cache"
	"log"
	"os"
)

func main() {
	mongo, err := cache.Mongo("mongodb://localhost")
	if err != nil {
		log.Fatal(err)
	}
	db, err := claimStorage.New("http://plus.kipris.or.kr/kipo-api/kipi/patUtiModInfoSearchSevice/getWordSearch", "http://plus.kipris.or.kr/kipo-api/kipi/patUtiModInfoSearchSevice/getBibliographyDetailInfoSearch", os.Getenv("KIPRIS"), mongo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(db)
	s := New(80, db)
	defer s.Close()
	err = s.Start()
	log.Fatal(err)
}
