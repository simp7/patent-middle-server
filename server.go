package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/logger"
	"github.com/simp7/patent-middle-server/model"
	"github.com/simp7/patent-middle-server/model/formula"
	"io"
)

type server struct {
	*gin.Engine
	Storage
	*logger.Logger
	port string
	fs   FileSystem
}

func New(port int, storage Storage, fs FileSystem, lg *logger.Logger) *server {

	s := new(server)

	s.Engine = gin.Default()
	s.Logger = lg
	s.Storage = storage
	s.port = fmt.Sprintf(":%d", port)
	s.fs = fs

	s.Info("server has been initialized")

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

	s.Info("start search")
	claims := s.GetClaims(input)

	if claims.IsEmpty() {
		_, _ = c.Writer.WriteString("<h1>Cannot find matching patent.</h1>")
		return
	}

	s.Info("create file")
	if err := s.fs.SaveCSVFile(claims); err != nil {
		s.Fatal(err)
	}

	//defer func() {
	//	if err := s.fs.RemoveCSVFile(claims); err != nil {
	//		s.Error(err)
	//	}
	//}()

	s.Info("perform NLP")
	if data, err := s.performNLP(country, claims); err == nil {
		s.writeResult(c.Writer, data)
		return
	}

	c.Writer.WriteHeader(500)

}

func (s *server) writeResult(writer io.Writer, data []byte) {
	if _, err := writer.Write(data); err != nil {
		s.Error(err)
		return
	}
	s.Info("search finished successfully")
}

func (s *server) Hello(c *gin.Context) {
	if _, err := c.Writer.WriteString("<h1>Server is Available</h1>"); err != nil {
		s.Error(err)
	}
}

func (s *server) performNLP(country string, group *model.CSVGroup) ([]byte, error) {

	args := make([]string, 2)
	args[0] = group.FileName
	args[1] = "10"

	switch country {
	case "KR":
		s.Info("select LDA")
		return s.fs.LDA(args...)
	case "LSA":
		s.Info("select LSA")
		return s.fs.LSA(args...)
	case "US":
		fallthrough
	default:
		s.Info("select Word2vec")
		keywords := formula.Interpret(group.SearchWord).Keywords()
		args = append(args, keywords...)
		return s.fs.Word2vec(args...)
	}

}
