package claimStorage

import (
	"encoding/xml"
	"github.com/google/logger"
	"github.com/simp7/patent-middle-server/model"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"
)

var instance *kipris
var once sync.Once

type kipris struct {
	*http.Client
	*logger.Logger
	cache     Cache
	apiKey    string
	SearchURL string
	ClaimURL  string
}

func New(searchURL string, claimURL string, apiKey string, cacheDB Cache) (*kipris, error) {

	var err error

	once.Do(func() {

		instance = &kipris{
			&http.Client{},
			logger.Init("server", true, false, os.Stdout),
			cacheDB,
			apiKey,
			searchURL,
			claimURL,
		}

	})

	return instance, err

}

func (k *kipris) GetClaims(input string) ([]model.CSVUnit, error) {

	numbers, err := k.searchNumbers(input)
	if err != nil {
		return nil, err
	}

	return k.searchClaims(numbers)

}

func (k *kipris) searchNumbers(input string) (result []string, err error) {

	rowSize := 20
	var once sync.Once

	request, err := http.NewRequest("GET", k.SearchURL, nil)
	if err != nil {
		return nil, err
	}

	q := request.URL.Query()
	q.Add("ServiceKey", k.apiKey)
	q.Add("word", input)
	q.Add("numOfRows", strconv.Itoa(rowSize))

	request.URL.RawQuery = q.Encode()

	response, err := k.Do(request)
	if err != nil {
		return nil, err
	}

	total, err := k.getTotal(response.Body)
	lastPage := total/rowSize + 1

	wg := sync.WaitGroup{}
	wg.Add(lastPage)

	for i := 1; i <= lastPage; i++ {
		go func(page int) {

			defer wg.Done()

			q.Set("pageNo", strconv.Itoa(page))
			request.URL.RawQuery = q.Encode()

			response, _ = k.Do(request)
			defer func() {
				once.Do(func() {
					if err = response.Body.Close(); err != nil {
						k.Error(err)
					}
				})
			}()

			result = append(result, k.searchNumber(response.Body)...)

		}(i)
	}

	wg.Wait()

	err = response.Body.Close()

	return result, err

}

func (k *kipris) getTotal(body io.Reader) (int, error) {

	var searchResult SearchResult
	err := xml.NewDecoder(body).Decode(&searchResult)

	if err != nil {
		return 0, err
	}

	return strconv.Atoi(searchResult.Count.TotalCount)

}

func (k *kipris) searchNumber(body io.Reader) (result []string) {

	var searchResult SearchResult

	err := xml.NewDecoder(body).Decode(&searchResult)
	if err != nil {
		return nil
	}

	items := searchResult.Body.Items.Item
	result = make([]string, len(items))

	for i, item := range items {
		result[i] = item.ApplicationNumber
	}

	return

}

func (k *kipris) searchClaims(numbers []string) (result []model.CSVUnit, err error) {

	request, err := http.NewRequest("GET", k.ClaimURL, nil)

	if err != nil {
		return nil, err
	}

	q := request.URL.Query()
	q.Add("ServiceKey", k.apiKey)

	wg := sync.WaitGroup{}
	wg.Add(len(numbers))

	for _, v := range numbers {
		go func(number string) {

			var claim model.CSVUnit

			defer wg.Done()

			tuple, ok := k.cache.Find(number)
			if ok {

				claim = tuple.Process()

			} else {

				q.Set("applicationNumber", number)
				request.URL.RawQuery = q.Encode()

				response, err := k.Do(request)
				if err != nil {
					k.Error(err)
					return
				}

				claim = k.restClaim(response.Body)

			}

			result = append(result, claim)

		}(v)
	}

	wg.Wait()

	return

}

func (k *kipris) restClaim(body io.Reader) model.CSVUnit {

	var searchResult ClaimResult
	k.Error(xml.NewDecoder(body).Decode(&searchResult))

	applicationNumber := searchResult.Body.Item.BiblioSummaryInfoArray.BiblioSummaryInfo.ApplicationNumber
	title := searchResult.Body.Item.BiblioSummaryInfoArray.BiblioSummaryInfo.InventionTitle
	claims := searchResult.Body.Item.ClaimInfoArray.ClaimInfo

	result := make([]string, len(claims))

	for i, claim := range claims {
		result[i] = claim.Claim
	}

	tuple := ClaimTuple{applicationNumber, title, result}
	k.Error(k.cache.Register(tuple))

	return tuple.Process()

}
