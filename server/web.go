package server

import (
	"fmt"
	"github.com/google/logger"
	"net/http"
	"os"
)

type web struct {
	http.Server
	*logger.Logger
}

func NewWeb(port int) *web {

	w := new(web)
	w.Logger = logger.Init("server", true, false, os.Stdout)
	w.Infof("Initialize logger")

	address := fmt.Sprintf(":%d", port)
	w.Server = http.Server{Addr: address}

	return w

}

func (w *web) Start() error {

	w.Info("Start Server")

	http.HandleFunc("/", w.Welcome)
	http.HandleFunc("/search", w.Search)

	return w.ListenAndServe()

}

func (w *web) Welcome(writer http.ResponseWriter, _ *http.Request) {
	writer.Write([]byte("<h1>Hello, world!</h1>"))
}

func (w *web) Search(writer http.ResponseWriter, request *http.Request) {

	country := unwrap(request, "country")
	formula := unwrap(request, "formula")

	w.Infof("GET %s in NLP of %s", formula, country)
	selected := w.selectNLP(country)
	w.Info("start search")

	for {
		w.Infof("Current formula is %s", formula)

		if formula == selected.Process(formula) {
			w.Infof("Final result: %s", formula)
			break
		}
		w.Infof("%s", formula)

	}

	_, err := writer.Write([]byte(formula))
	if err != nil {
		w.Error(err)
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

func (w *web) selectNLP(country string) *nlp {

	switch country {
	case "KR":
		w.Info("Select Korean")
		return Korean()
	case "US":
		fallthrough
	default:
		w.Info("Select English")
		return English()
	}

}
