package main

import (
	"github.com/simp7/patent-middle-server/server"
)

func main() {

	s := server.NewWeb(80)
	defer s.Server.Close()
	s.Start()

}
