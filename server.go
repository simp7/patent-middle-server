package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/logger"
	"github.com/simp7/patent-middle-server/model/formula"
	"os"
)

type server struct {
	*gin.Engine
	Storage
	*logger.Logger
	port string
}

func New(port int, storage Storage, lg *logger.Logger) *server {

	s := new(server)

	s.Engine = gin.Default()
	s.Logger = lg

	s.Infof("Finish Initializing Logger")

	s.Storage = storage

	s.port = fmt.Sprintf(":%d", port)

	return s

}

func (s *server) Close() {
	s.Logger.Close()
}

func (s *server) Start() error {
	s.GET("/:country/:formula", s.Search)
	s.GET("/", s.Hello)
	s.Info("start server")
	return s.Run(s.port)
}

func (s *server) Search(c *gin.Context) {

	country := c.Param("country")
	input := c.Param("formula")

	s.Infof("GET %s in NLP of %s", input, country)
	selected := s.selectNLP(country)

	s.Info("start search")
	claims := s.GetClaims(input)

	s.Info("create file")
	file, err := claims.File()
	if err != nil {
		s.Fatal(err)
	}

	defer func() {
		if err = os.Remove(file.Name()); err != nil {
			s.Error(err)
		}
	}()

	s.Info("perform NLP")
	data, err := selected.Process(file.Name(), formula.Interpret(input).KeyWords()...)

	if err != nil {
		s.Error(err)
		c.Writer.WriteHeader(500)
		return
	}

	if _, err = c.Writer.Write(data); err != nil {
		s.Error(err)
	}

}

func (s *server) Hello(c *gin.Context) {
	if _, err := c.Writer.WriteString("<h1>Server is Available</h1>"); err != nil {
		s.Error(err)
	}
}

func (s *server) selectNLP(country string) NLP {

	switch country {
	case "KR":
		s.Info("select LDA")
		return LDA()
	case "US":
		fallthrough
	default:
		s.Info("select Word2vec")
		return Word2vec()
	}

}
