package main

import (
	"fmt"
	"github.com/simp7/patent-middle-server/claimDB"
	"log"
	"os"
)

func main() {
	db, err := claimDB.New("http://plus.kipris.or.kr/kipo-api/kipi/patUtiModInfoSearchSevice/getWordSearch", "http://plus.kipris.or.kr/kipo-api/kipi/patUtiModInfoSearchSevice/getBibliographyDetailInfoSearch", os.Getenv("KIPRIS"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(db)
	s := New(80, db)
	defer s.Close()
	s.Start()
}
