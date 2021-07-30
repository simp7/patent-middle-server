package claimDB

import (
	"encoding/xml"
	"io"
	"net/http"
	"strconv"
	"sync"
)

type kipo struct {
	*http.Client
	apiKey string
	url    string
}

func Kipo(apiKey string) *kipo {
	return &kipo{
		&http.Client{},
		apiKey,
		"http://kipo-api.kipi.or.kr/openapi/service/patUtiModInfoSearchSevice/getAdvancedSearch",
	}
}

func (k *kipo) searchNumbers(input string) (result []string, err error) {

	rowSize := 500
	totalCount := 1
	var response *http.Response

	searchResult := make([]SearchResult, 0)

	request, err := http.NewRequest("GET", k.url, nil)
	if err != nil {
		return nil, err
	}

	q := request.URL.Query()
	q.Add("ServiceKey", k.apiKey)
	q.Add("inventionTitle", input)
	q.Add("numOfRows", strconv.Itoa(rowSize))

	for page := 1; (totalCount-1)/rowSize <= page-1; page++ {
		var tmpResult SearchResult
		q.Add("pageNo", strconv.Itoa(page))
		request.URL.RawQuery = q.Encode()
		response, err = k.Do(request)
		if err != nil {
			return nil, err
		}
		tmpResp, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		xml.Unmarshal(tmpResp, &tmpResult)

		//for _, item := range tmpResult.Body.Items.Item {
		//	//tokenize
		//}
		searchResult = append(searchResult, tmpResult)
	}
	defer response.Body.Close()

	wg := sync.WaitGroup{}

	wg.Wait()

	return

}
