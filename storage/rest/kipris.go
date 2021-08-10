package rest

import (
	"encoding/xml"
	"github.com/google/logger"
	"github.com/simp7/patent-middle-server/storage"
	"io"
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

func New(config Config) *kipris {
	key := config.Key
	if key == "" {
		key = os.Getenv("KIPRIS")
	}
	return &kipris{
		&http.Client{},
		logger.Init("server", true, false, os.Stdout),
		500,
		key,
		config.WordURL,
		config.ClaimURL,
	}
}

func (k *kipris) GetClaims(number string) storage.ClaimTuple {

	k.Info("getting claims of patent : " + number)

	request, err := http.NewRequest("GET", k.ClaimURL, nil)
	k.check(err)

	q := request.URL.Query()
	q.Add("ServiceKey", k.apiKey)
	q.Add("applicationNumber", number)
	request.URL.RawQuery = q.Encode()

	k.Info("send " + request.URL.RawQuery)
	response, err := k.Do(request)
	if err != nil {
		k.Error(err)
	}

	return k.processClaim(response.Body)

}

func (k *kipris) GetNumbers(input string) chan chan string {

	k.Info("getting application numbers by searching " + input)
	outCh := make(chan chan string)

	request, err := http.NewRequest("GET", k.SearchURL, nil)
	k.check(err)

	q := request.URL.Query()
	q.Add("ServiceKey", k.apiKey)
	q.Add("word", input)
	q.Add("numOfRows", strconv.Itoa(k.pageRow))

	request.URL.RawQuery = q.Encode()

	response, err := k.Do(request)
	k.check(err)

	total, err := k.getTotal(response.Body)
	if err != nil {
		return outCh
	}

	lastPage := (total-1)/k.pageRow + 1

	outCh = make(chan chan string, lastPage)

	wg := sync.WaitGroup{}
	wg.Add(lastPage)

	for i := 1; i <= lastPage; i++ {
		go func(page int) {

			q.Set("pageNo", strconv.Itoa(page))
			request.URL.RawQuery = q.Encode()

			response, _ = k.Do(request)

			outCh <- k.getNumberByPage(response.Body)

			wg.Done()

		}(i)
	}

	wg.Wait()
	k.Info("close channel in GetNumbers")
	close(outCh)

	defer func() {
		k.check(response.Body.Close())
	}()

	return outCh

}

func (k *kipris) getTotal(body io.Reader) (int, error) {

	var searchResult storage.SearchResult
	err := xml.NewDecoder(body).Decode(&searchResult)

	if err != nil {
		return 0, err
	}

	return strconv.Atoi(searchResult.Count.TotalCount)

}

func (k *kipris) getNumberByPage(body io.Reader) chan string {

	outCh := make(chan string, k.pageRow)
	var searchResult storage.SearchResult

	err := xml.NewDecoder(body).Decode(&searchResult)
	if err != nil {
		return nil
	}

	items := searchResult.Body.Items.Item
	var wg sync.WaitGroup

	wg.Add(len(items))
	a := len(items)
	for _, item := range items {
		go func(number string) {
			k.Info("number " + number + " founded")
			outCh <- number
			wg.Done()
			a--
		}(item.ApplicationNumber)
	}

	go func() {
		wg.Wait()
		k.Info("close channel in getNumberByPage")
		close(outCh)
	}()

	return outCh

}

func (k *kipris) processClaim(body io.Reader) storage.ClaimTuple {

	var searchResult storage.ClaimResult
	k.check(xml.NewDecoder(body).Decode(&searchResult))

	applicationNumber := searchResult.Body.Item.BiblioSummaryInfoArray.BiblioSummaryInfo.ApplicationNumber
	title := searchResult.Body.Item.BiblioSummaryInfoArray.BiblioSummaryInfo.InventionTitle
	claims := searchResult.Body.Item.ClaimInfoArray.ClaimInfo

	result := make([]string, len(claims))

	for i, claim := range claims {
		result[i] = claim.Claim
	}

	return storage.ClaimTuple{ApplicationNumber: applicationNumber, Name: title, Claims: result}

}

func (k *kipris) check(err error) {
	if err != nil {
		k.Error(err)
	}
}
