package server

import "github.com/kataras/golog"

type Server struct {
	*web
	kNLP *nlp
	eNLP *nlp
}

func New(level golog.Level, port int) *Server {

	s := new(Server)
	s.web = newWeb(level, port)

	return s

}

func (s *Server) a() {

}
