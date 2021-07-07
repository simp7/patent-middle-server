package server

import (
	"fmt"
	"github.com/google/logger"
	"github.com/simp7/patent-middle-server/server/claimDB"
	"net/http"
	"os"
)

type server struct {
	http.Server
	ClaimDB
	*logger.Logger
}

func New(port int) *server {

	w := new(server)
	w.Logger = logger.Init("server", true, false, os.Stdout)
	w.Infof("Initialize logger")

	address := fmt.Sprintf(":%d", port)
	w.Server = http.Server{Addr: address}

	w.ClaimDB = claimDB.New(os.Getenv("KIPRIS"))

	return w

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
	writer.Write([]byte("<h1>Hello, world!</h1>"))
}

func (s *server) Search(writer http.ResponseWriter, request *http.Request) {

	country := unwrap(request, "country")
	input := unwrap(request, "formula")

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

func (s *server) processNLP(instance nlp, input string) ([]unit, error) {
	_, err := instance.Process(input)
	if err != nil {
		return nil, err
	}
	return ProcessCSV(os.Stdin)
}

func unwrap(request *http.Request, key string) string {
	result := ""
	if value := request.URL.Query()[key]; len(value) != 0 {
		result = value[0]
	}
	return result
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
