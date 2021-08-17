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

func New(port int, storage Storage) *server {

	s := new(server)

	s.Engine = gin.Default()

	s.Logger = logger.Init("server", true, false, os.Stdout)
	s.Infof("Finish Initializing Logger")

	s.Storage = storage

	s.port = fmt.Sprintf(":%d", port)

	return s

}

func (s *server) Close() {
	s.Logger.Close()
}

func (s *server) Start() error {

	s.Info("start server")

	s.GET("/", s.Welcome)
	s.GET("/claims/:country/:formula", s.Search)

	return s.Run(s.port)

}

func (s *server) Welcome(c *gin.Context) {
	_, err := c.Writer.Write([]byte("<h1>Hello, world!</h1>"))
	if err != nil {
		s.Error(err)
	}
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
		err = os.Remove(file.Name())
		if err != nil {
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

	_, err = c.Writer.Write(data)
	if err != nil {
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
