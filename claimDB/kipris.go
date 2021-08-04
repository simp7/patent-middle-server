package claimDB

import (
	"context"
	"encoding/xml"
	"github.com/simp7/patent-middle-server/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var instance *kipris
var once sync.Once
var dbServer = "localhost" // DB 서버 위치 지정

type kipris struct {
	collection *mongo.Collection
	*http.Client
	apiKey    string
	SearchURL string
	ClaimURL  string
}

func New(apiKey string) *kipris {

	once.Do(func() {

		var client *mongo.Client

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbServer))
		if err != nil {
			return
		}

		db := client.Database("Patent")
		collection := db.Collection("claim")

		instance = &kipris{
			collection,
			&http.Client{},
			apiKey,
			"http://plus.kipris.or.kr/kipo-api/kipi/patUtiModInfoSearchSevice/getWordSearch",
			"http://plus.kipris.or.kr/kipo-api/kipi/patUtiModInfoSearchSevice/getBibliographyDetailInfoSearch",
		}

	})

	return instance

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

			var claim model.CSVUnit

			dbResult := k.collection.FindOne(context.TODO(), bson.D{{"applicationNumber", number}})
			if dbResult.Err() != nil {

				var tuple ClaimTuple
				dbResult.Decode(&tuple)
				claim = tuple.Process()

			} else {

				q.Set("applicationNumber", number)
				request.URL.RawQuery = q.Encode()

				response, _ := k.Do(request)
				defer response.Body.Close()
				claim = k.searchClaim(response.Body)

			}

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

	applicationNumber := searchResult.Body.Items.BiblioSummaryInfoArray.BiblioSummaryInfo.ApplicationNumber
	title := searchResult.Body.Items.BiblioSummaryInfoArray.BiblioSummaryInfo.InventionTitle
	claims := searchResult.Body.Items.ClaimInfoArray.ClaimInfo
	result := make([]string, len(claims))

	for i, claim := range claims {
		result[i] = claim.Claim
	}

	tuple := ClaimTuple{applicationNumber, title, result}
	k.putToDB(tuple)

	return tuple.Process()

}

func (k *kipris) putToDB(tuple ClaimTuple) error {
	_, err := k.collection.InsertOne(context.TODO(), tuple)
	return err
}
