package rest

import (
	"encoding/xml"
	"errors"
	"github.com/google/logger"
	"github.com/simp7/patent-middle-server/storage"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
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

func (k *kipris) GetClaims(number string) storage.ClaimTuple {

	k.Info("getting claims of patent : " + number)

	request, err := http.NewRequest("GET", k.ClaimURL, nil)
	k.check(err)

	q := request.URL.Query()
	q.Add("ServiceKey", k.apiKey)
	q.Add("applicationNumber", number)
	request.URL.RawQuery = q.Encode()

	k.Info("send " + request.URL.String())
	response, err := k.Do(request)
	k.check(err)

	return k.processClaim(response.Body)

}

func (k *kipris) GetNumbers(input string) chan chan string {

	var totalPage int

	k.Info("getting application numbers by searching " + input)
	outCh := make(chan chan string)

	request, err := http.NewRequest("GET", k.SearchURL, nil)
	k.check(err)

	q := request.URL.Query()
	q.Add("ServiceKey", k.apiKey)
	q.Add("word", input)
	q.Add("numOfRows", strconv.Itoa(k.pageRow))

	request.URL.RawQuery = q.Encode()

	k.Info("send " + request.URL.RawQuery)
	response, err := k.Do(request)
	k.check(err)

	if totalPage, err = k.getTotal(response.Body); err != nil {
		k.Error(err)
		close(outCh)
		return outCh
	}

	lastPage := (totalPage-1)/k.pageRow + 1

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

	defer k.check(response.Body.Close())

	return outCh

}

func (k *kipris) getTotal(body io.Reader) (int, error) {

	var searchResult storage.SearchResult
	if err := xml.NewDecoder(body).Decode(&searchResult); err != nil {
		return 0, errors.New("failed getting data from kipris -- check your api key")
	}

	k.Info("getting total pages of application numbers")
	return strconv.Atoi(searchResult.Count.TotalCount)

}

func (k *kipris) getNumberByPage(body io.Reader) chan string {

	outCh := make(chan string, k.pageRow)
	var searchResult storage.SearchResult

	if err := xml.NewDecoder(body).Decode(&searchResult); err != nil {
		k.Error(err)
		return nil
	}

	items := searchResult.Body.Items.Item
	var wg sync.WaitGroup

	wg.Add(len(items))
	a := len(items)
	for _, item := range items {
		go func(number string) {
			k.Info("getting " + number)
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

	title = strings.TrimSpace(title)

	result := make([]string, len(claims))
	for i, claim := range claims {
		result[i] = strings.TrimSpace(claim.Claim)
	}

	return k.trim(applicationNumber, title, result)

}

func (k *kipris) trim(applicationNumber string, name string, claims []string) storage.ClaimTuple {

	resultNumber := strings.Join(strings.Split(applicationNumber, "-"), "")
	resultName := strings.TrimSpace(name)

	return storage.ClaimTuple{ApplicationNumber: resultNumber, Name: resultName, Claims: claims}

}

func (k *kipris) check(err error) {
	if err != nil {
		k.Error(err)
	}
}
