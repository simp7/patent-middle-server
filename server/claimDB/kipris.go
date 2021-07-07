package claimDB

import (
	"encoding/csv"
	"encoding/xml"
	"io"
	"net/http"
	"sync"
)

type kipris struct {
	*http.Client
	key       string
	SearchURL string
	ClaimURL  string
}

func New(key string) *kipris {
	return &kipris{
		&http.Client{},
		key,
		"http://kipo-api.kipi.or.kr/openapi/service/patUtiModInfoSearchSevice/getAdvancedSearch",
		"http://kipo-api.kipi.or.kr/openapi/service/patUtiModInfoSearchSevice/getBibliographyDetailInfoSearch",
	}
}

func (k *kipris) GetClaims(input string) (*csv.Reader, error) {

	numbers, err := k.searchNo(input)
	if err != nil {
		return nil, err
	}

	wg := sync.WaitGroup{}

	for _, v := range numbers {
		wg.Add(1)
		go func() {
			k.searchClaim(v) //TODO: 에러 처리 구문 및 claim 가공

			wg.Done()
		}()
	}
	wg.Wait()

	return nil, nil

}

func (k *kipris) searchNo(input string) ([]string, error) {

	request, err := http.NewRequest("GET", k.SearchURL, nil)
	if err != nil {
		return nil, err
	}

	q := request.URL.Query()
	q.Add("ServiceKey", k.key)
	q.Add("inventionTitle", input)
	q.Add("lastvalue", "R")
	request.URL.RawQuery = q.Encode()

	return k.solve(request, "applicationNumber")

}

func (k *kipris) searchClaim(patentNo string) ([]string, error) {

	request, err := http.NewRequest("GET", k.ClaimURL, nil)
	if err != nil {
		return nil, err
	}

	q := request.URL.Query()
	q.Add("ServiceKey", k.key)
	q.Add("applicationNumber", patentNo)
	request.URL.RawQuery = q.Encode()

	return k.solve(request, "claim")

}

func (k *kipris) solve(request *http.Request, parseTarget string) ([]string, error) {

	response, err := k.Do(request)
	if err != nil {
		return nil, err
	}

	return parse(response.Body, parseTarget), nil

}

func parse(reader io.Reader, element string) (result []string) {

	result = make([]string, 0)
	parser := xml.NewDecoder(reader)

	for {

		token, err := parser.Token()
		if err != nil {
			break
		}

		switch t := token.(type) {
		case xml.StartElement:
			if t.Name.Local == element {
				token, err = parser.Token()
				result = append(result, extractData(token))
			}
		}

	}

	return

}

func extractData(token xml.Token) string {
	switch t := token.(type) {
	case xml.CharData:
		return string(t)
	}
	return ""
}
