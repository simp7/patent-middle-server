package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/logger"
	"github.com/simp7/patent-middle-server/model/formula"
	"github.com/simp7/patent-middle-server/nlp"
	"os"
)

type server struct {
	*gin.Engine
	ClaimDB
	*logger.Logger
	port string
}

func New(port int, claimDB ClaimDB) *server {

	s := new(server)

	s.Engine = gin.Default()

	s.Logger = logger.Init("server", true, false, os.Stdout)
	s.Infof("Finish Initializing logger")

	s.ClaimDB = claimDB

	s.port = fmt.Sprintf(":%d", port)

	return s

}

func (s *server) Close() {
	s.Logger.Close()
}

func (s *server) Start() error {

	s.Info("Start Server")

	s.GET("/", s.Welcome)
	s.GET("/search/:country/:formula", s.Search)

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
	_, err := s.GetClaims(input)
	if err != nil {
		s.Error(err)
		c.Writer.WriteHeader(500)
		return
	}

	s.Info("perform NLP")
	data, err := selected.Process(input+".csv", formula.Interpret(input).KeyWords()...)

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
		s.Info("Select LDA")
		return nlp.LDA()
	case "US":
		fallthrough
	default:
		s.Info("Select Word2vec")
		return nlp.Word2vec()
	}

}
