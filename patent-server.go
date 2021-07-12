package main

import (
	"github.com/simp7/patent-middle-server/claimDB"
	"os"
)

func main() {
	s := New(80, claimDB.New(os.Getenv("KIPRIS")))
	defer s.Close()
	s.Start()
}
