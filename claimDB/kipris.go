package claimDB

import (
	"encoding/xml"
	"github.com/simp7/patent-middle-server/model"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type kipris struct {
	*http.Client
	apiKey    string
	SearchURL string
	ClaimURL  string
}

func New(apiKey string) *kipris {
	return &kipris{
		&http.Client{},
		apiKey,
		"http://plus.kipris.or.kr/kipo-api/kipi/patUtiModInfoSearchSevice/getWordSearch",
		"http://plus.kipris.or.kr/kipo-api/kipi/patUtiModInfoSearchSevice/getBibliographyDetailInfoSearch",
	}
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

			q.Set("pageNo", strconv.Itoa(page))
			request.URL.RawQuery = q.Encode()

			response, _ = k.Do(request)
			defer response.Body.Close()

			wg.Done()

		}(i)
	}

	wg.Wait()

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
	xml.NewDecoder(body).Decode(&searchResult)

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

			q.Set("applicationNumber", number)
			request.URL.RawQuery = q.Encode()

			response, _ := k.Do(request)
			defer response.Body.Close()
			claim := k.searchClaim(response.Body)

			result = append(result, claim)

			wg.Done()

		}(v)
	}

	wg.Wait()

	return

}

func (k *kipris) searchClaim(body io.Reader) model.CSVUnit {

	var searchResult ClaimResult
	xml.NewDecoder(body).Decode(&searchResult)

	claims := searchResult.Body.Items.ClaimInfoArray.ClaimInfo
	result := make([]string, len(claims))
	for i, claim := range claims {
		result[i] = claim.Claim
	}

	return tokenize(searchResult.Body.Items.BiblioSummaryInfoArray.BiblioSummaryInfo.InventionTitle, result)

}

func tokenize(title string, claims []string) model.CSVUnit {
	return model.CSVUnit{Key: title, Value: strings.Join(claims, "\n")}
}
