package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/logger"
	"github.com/simp7/patent-middle-server/customcsv"
	"github.com/simp7/patent-middle-server/model"
	"github.com/simp7/patent-middle-server/nlp"
	"os"
	"strings"
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

	//TODO: 청구항 csv 파일을 NLP로 전달

	s.Info("perform NLP")
	_, err = s.processNLP(selected, input) //TODO: NLP를 통해 얻은 단어-유사도 클라이언트에게 전달
	//if err != nil {
	//	s.Error(err)
	//	c.Writer.WriteHeader(500)
	//	return
	//}

	_, err = c.Writer.Write([]byte(input))
	if err != nil {
		s.Error(err)
	}

}

func (s *server) processNLP(instance NLP, input string) ([]model.CSVUnit, error) {

	s.Infof("Process %s in NLP", input)
	result, err := instance.Process(input)
	if err != nil {
		return nil, err
	}

	reader := strings.NewReader(result)

	return customcsv.Parser(reader).Parse()

}

func (s *server) selectNLP(country string) NLP {

	switch country {
	case "KR":
		s.Info("Select Korean")
		return nlp.Korean()
	case "US":
		fallthrough
	default:
		s.Info("Select English")
		return nlp.English()
	}

}
