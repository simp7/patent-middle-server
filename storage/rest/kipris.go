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
	apiKey    string
	SearchURL string
	ClaimURL  string
}

func New(searchURL string, claimURL string, apiKey string) *kipris {
	return &kipris{
		&http.Client{},
		logger.Init("server", true, false, os.Stdout),
		apiKey,
		searchURL,
		claimURL,
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

	response, err := k.Do(request)
	if err != nil {
		k.Error(err)
	}

	return k.processClaim(response.Body)

}

func (k *kipris) GetNumbers(input string) chan string {

	k.Info("getting application numbers by searching " + input)
	outCh := make(chan string)
	defer close(outCh)

	rowSize := 500

	request, err := http.NewRequest("GET", k.SearchURL, nil)
	k.check(err)

	q := request.URL.Query()
	q.Add("ServiceKey", k.apiKey)
	q.Add("word", input)
	q.Add("numOfRows", strconv.Itoa(rowSize))

	request.URL.RawQuery = q.Encode()

	response, err := k.Do(request)
	k.check(err)

	total, err := k.getTotal(response.Body)
	k.check(err)
	lastPage := total/rowSize + 1

	wg := sync.WaitGroup{}
	wg.Add(lastPage)

	for i := 1; i <= lastPage; i++ {
		go func(page int) {

			defer wg.Done()

			q.Set("pageNo", strconv.Itoa(page))
			request.URL.RawQuery = q.Encode()

			response, _ = k.Do(request)

			outCh <- <-k.getNumberByPage(response.Body)

		}(i)
	}

	wg.Wait()

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

	outCh := make(chan string)
	var searchResult storage.SearchResult

	err := xml.NewDecoder(body).Decode(&searchResult)
	if err != nil {
		return nil
	}

	items := searchResult.Body.Items.Item

	for _, item := range items {
		go func(number string) {
			outCh <- number
		}(item.ApplicationNumber)
	}

	return outCh

}

func (k *kipris) processClaim(body io.Reader) storage.ClaimTuple {

	var searchResult storage.ClaimResult
	k.Error(xml.NewDecoder(body).Decode(&searchResult))

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
