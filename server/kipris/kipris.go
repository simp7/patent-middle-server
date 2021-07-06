package kipris

import (
	"encoding/csv"
	"encoding/xml"
	"net/http"
)

type kipris struct {
	*http.Client
	key       string
	SearchURL string
	ClaimURL  string
}

func New(key string) *kipris {
	return &kipris{&http.Client{}, key, "http://plus.kipris.or.kr/kipo-api/kipi/patUtiModInfoSearchSevice/getAdvancedSearch", "http://plus.kipris.or.kr/kipo-api/kipi/patUtiModInfoSearchSevice/getBibliographyDetailInfoSearch"}
}

func (k *kipris) GetClaims(input string) (*csv.Reader, error) {

	return nil, nil

}

func (k *kipris) searchNo(input string) ([]string, error) {

	req, err := http.NewRequest("GET", k.ClaimURL, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("ServiceKey", k.key)
	q.Add("inventionTitle", input)
	q.Add("lastvalue", "R")
	req.URL.RawQuery = q.Encode()

	resp, err := k.Do(req)

	parser := xml.NewDecoder(resp.Body)
	for {
		_, err := parser.Token()
		if err != nil {
			break
		}
	}

	return nil, nil

}

func (k *kipris) searchClaims() error {
	return nil
}
