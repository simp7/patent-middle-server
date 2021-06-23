package main

import (
	"github.com/kataras/golog"
	"github.com/simp7/patent-middle-server/server"
)

func main() {

	server.New(golog.DebugLevel, 443)

}
