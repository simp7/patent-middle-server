package main

import (
	"github.com/simp7/patent-middle-server/server"
	"github.com/simp7/patent-middle-server/server/claimDB"
	"os"
)

func main() {
	s := server.New(80, claimDB.New(os.Getenv("KIPRIS")))
	defer s.Close()
	s.Start()
}
