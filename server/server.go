package server

import (
	"fmt"
	"github.com/google/logger"
	"github.com/simp7/patent-middle-server/model"
	"github.com/simp7/patent-middle-server/server/csvParser"
	"net/http"
	"os"
	"strings"
)

type server struct {
	http.Server
	ClaimDB
	*logger.Logger
}

func New(port int, claimDB ClaimDB) *server {

	s := new(server)
	s.Logger = logger.Init("server", true, false, os.Stdout)
	s.Infof("Initialize logger")

	address := fmt.Sprintf(":%d", port)
	s.Server = http.Server{Addr: address}

	s.ClaimDB = claimDB

	return s

}

func (s *server) Close() error {

	err := s.Server.Close()
	if err != nil {
		s.Error(err)
	}

	s.Logger.Close()
	return err

}

func (s *server) Start() error {

	s.Info("Start Server")

	http.HandleFunc("/", s.Welcome)
	http.HandleFunc("/search", s.Search)

	return s.ListenAndServe()

}

func (s *server) Welcome(writer http.ResponseWriter, _ *http.Request) {
	_, err := writer.Write([]byte("<h1>Hello, world!</h1>"))
	if err != nil {
		s.Error(err)
	}
}

func (s *server) Search(writer http.ResponseWriter, request *http.Request) {

	country := s.unwrap(request, "country")
	input := s.unwrap(request, "formula")

	s.Infof("GET %s in NLP of %s", input, country)
	selected := s.selectNLP(country)

	s.Info("start search")
	s.GetClaims(input)

	s.Info("perform NLP")
	_, err := s.processNLP(selected, input)
	if err != nil {
		s.Error(err)
		writer.WriteHeader(500)
	}

	_, err = writer.Write([]byte(input))
	if err != nil {
		s.Error(err)
	}

}

func (s *server) processNLP(instance nlp, input string) ([]model.CSVUnit, error) {

	s.Infof("Process %s in NLP", input)
	result, err := instance.Process(input)
	if err != nil {
		return nil, err
	}

	reader := strings.NewReader(result)

	return csvParser.Parser(reader).Parse()

}

func (s *server) unwrap(request *http.Request, key string) (result string) {
	if value := request.URL.Query()[key]; len(value) != 0 {
		result = value[0]
	}
	s.Infof("Find %s: %s", key, result)
	return
}

func (s *server) selectNLP(country string) nlp {

	switch country {
	case "KR":
		s.Info("Select Korean")
		return Korean()
	case "US":
		fallthrough
	default:
		s.Info("Select English")
		return English()
	}

}
