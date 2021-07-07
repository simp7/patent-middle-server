package server

import (
	"fmt"
	"github.com/google/logger"
	"net/http"
	"os"
)

type server struct {
	http.Server
	*logger.Logger
}

func New(port int) *server {

	w := new(server)
	w.Logger = logger.Init("server", true, false, os.Stdout)
	w.Infof("Initialize logger")

	address := fmt.Sprintf(":%d", port)
	w.Server = http.Server{Addr: address}

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
	formula := unwrap(request, "formula")

	s.Infof("GET %s in NLP of %s", formula, country)
	selected := s.selectNLP(country)

	s.Info("start search")

	for {

		s.Infof("Current formula is %s", formula)
		calculated, err := selected.Process(formula)

		if err != nil {
			s.Error(err)
			writer.WriteHeader(500)
			break
		}

		if formula == calculated {
			s.Infof("Final result: %s", formula)
			break
		}

	}

	_, err := writer.Write([]byte(formula))
	if err != nil {
		s.Error(err)
	}
	//TODO: KIPRIS에서 API를 통한 검색 및 출력

}

func unwrap(request *http.Request, key string) string {
	result := ""
	if value := request.URL.Query()[key]; len(value) != 0 {
		result = value[0]
	}
	return result
}

func (s *server) selectNLP(country string) *nlp {

	//TODO: reader 수정

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
