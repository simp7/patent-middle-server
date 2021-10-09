package rest

import (
	"encoding/xml"
	"errors"
	"github.com/google/logger"
	"github.com/simp7/patent-middle-server/storage"
	"net/http"
	"os"
	"strconv"
	"sync"
)

type kipris struct {
	*http.Client
	*logger.Logger
	pageRow   int
	apiKey    string
	SearchURL string
	ClaimURL  string
}

func New(config Config, lg *logger.Logger) *kipris {
	key := config.Key
	if key == "" {
		key = os.Getenv("KIPRIS")
	}
	return &kipris{
		&http.Client{},
		lg,
		500,
		key,
		config.WordURL,
		config.ClaimURL,
	}
}

func (k *kipris) GetClaims(number string) storage.Data {

	k.Info("getting claims of patent : " + number)

	queryResult, err := k.queryClaims(number)
	k.check(err)

	return k.processClaim(queryResult)

}

func (k *kipris) GetNumbers(input string, outCh chan<- string) {

	k.Info("find application numbers by searching " + input)

	var queryResult storage.SearchResult
	var err error
	var wg sync.WaitGroup

	getNumbers := func(page int) {

		defer wg.Done()

		queryResult, err = k.queryNumbers(input, page)
		k.check(err)

		queryResult.ApplicationNumbers(outCh)
		k.Info("got application numbers by page")

	}

	wg.Add(1)
	go getNumbers(1)

	total, err := queryResult.TotalPatent()
	k.check(err)

	k.Infof("total number of patent is %d", total)
	lastPage := (total-1)/k.pageRow + 1

	for i := 2; i <= lastPage; i++ {
		wg.Add(1)
		go getNumbers(i)
	}

	wg.Wait()

	k.Info("close channel in GetNumbers")
	close(outCh)

}

func (k *kipris) queryNumbers(input string, page int) (result storage.SearchResult, err error) {

	request, err := http.NewRequest("GET", k.SearchURL, nil)
	if err != nil {
		return
	}

	q := request.URL.Query()
	q.Add("ServiceKey", k.apiKey)
	q.Add("word", input)
	q.Add("numOfRows", strconv.Itoa(k.pageRow))
	q.Add("pageNo", strconv.Itoa(page))

	request.URL.RawQuery = q.Encode()

	k.Info("send " + request.URL.RawQuery)
	response, err := k.Do(request)
	if err != nil {
		return
	}

	if err = xml.NewDecoder(response.Body).Decode(&result); err != nil {
		err = errors.New("failed getting data from kipris -- check your api key")
	}

	return

}

func (k *kipris) queryClaims(number string) (result storage.ClaimResult, err error) {

	request, err := http.NewRequest("GET", k.ClaimURL, nil)
	if err != nil {
		return
	}

	q := request.URL.Query()
	q.Add("ServiceKey", k.apiKey)
	q.Add("applicationNumber", number)

	request.URL.RawQuery = q.Encode()

	k.Info("send " + request.URL.String())
	response, err := k.Do(request)
	if err != nil {
		return
	}

	err = xml.NewDecoder(response.Body).Decode(&result)

	return result, err

}

func (k *kipris) processClaim(claimData storage.ClaimResult) storage.Data {

	applicationNumber := claimData.ApplicationNumber()
	title := claimData.Title()
	claims := claimData.Claims()

	return storage.NewData(applicationNumber, title, claims)

}

func (k *kipris) check(err error) {
	if err != nil {
		k.Error(err)
	}
}
