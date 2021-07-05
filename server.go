package main

import (
	"github.com/simp7/patent-middle-server/server"
)

func main() {
	s := server.New(80)
	defer s.Close()
	s.Start()
}
