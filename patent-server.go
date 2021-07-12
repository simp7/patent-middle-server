package main

import (
	claimDB2 "github.com/simp7/patent-middle-server/claimDB"
	"os"
)

func main() {
	s := New(80, claimDB2.New(os.Getenv("KIPRIS")))
	defer s.Close()
	s.Start()
}
