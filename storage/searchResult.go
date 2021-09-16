package storage

import (
	"encoding/xml"
	"strconv"
	"sync"
)

type SearchResult struct {
	XMLName xml.Name `xml:"response"`
	Text    string   `xml:",chardata"`
	Header  struct {
		Text          string `xml:",chardata"`
		RequestMsgID  string `xml:"requestMsgID"`
		ResponseTime  string `xml:"responseTime"`
		ResponseMsgID string `xml:"responseMsgID"`
		SuccessYN     string `xml:"successYN"`
		ResultCode    string `xml:"resultCode"`
		ResultMsg     string `xml:"resultMsg"`
	} `xml:"header"`
	Body struct {
		Text  string `xml:",chardata"`
		Items struct {
			Text string `xml:",chardata"`
			Item []struct {
				Text              string `xml:",chardata"`
				ApplicantName     string `xml:"applicantName"`
				ApplicationDate   string `xml:"applicationDate"`
				ApplicationNumber string `xml:"applicationNumber"`
				AstrtCont         string `xml:"astrtCont"`
				BigDrawing        string `xml:"bigDrawing"`
				Drawing           string `xml:"drawing"`
				IndexNo           string `xml:"indexNo"`
				InventionTitle    string `xml:"inventionTitle"`
				IpcNumber         string `xml:"ipcNumber"`
				OpenDate          string `xml:"openDate"`
				OpenNumber        string `xml:"openNumber"`
				PublicationDate   string `xml:"publicationDate"`
				PublicationNumber string `xml:"publicationNumber"`
				RegisterDate      string `xml:"registerDate"`
				RegisterNumber    string `xml:"registerNumber"`
				RegisterStatus    string `xml:"registerStatus"`
			} `xml:"item"`
		} `xml:"items"`
	} `xml:"body"`
	Count struct {
		Text       string `xml:",chardata"`
		NumOfRows  string `xml:"numOfRows"`
		PageNo     string `xml:"pageNo"`
		TotalCount string `xml:"totalCount"`
	} `xml:"count"`
}

func (s SearchResult) TotalPage() int {
	result, _ := strconv.Atoi(s.Count.TotalCount)
	return result
}

func (s SearchResult) ApplicationNumbers(outCh chan<- string) {

	items := s.Body.Items.Item

	var wg sync.WaitGroup
	wg.Add(len(items))

	for _, item := range items {
		go func(number string) {
			outCh <- number
			wg.Done()
		}(item.ApplicationNumber)
	}
	wg.Wait()

}
