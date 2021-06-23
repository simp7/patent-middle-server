package server

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
)

type web struct {
	port string
	*iris.Application
}

func newWeb(level golog.Level, port int) *web {

	w := new(web)
	w.Logger().Level = level

	w.Application = iris.New()
	w.port = string(port)

	return w

}

func (w *web) Start() error {

	w.Get("/search/{patent}", w.Search)

	return w.Listen(w.port)

}

func (w *web) Search(ctx iris.Context) {

}
