package main

import (
	"github.com/simp7/patent-middle-server/storage"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"sync"
	"testing"
	"time"
)

var (
	testServer = New(80, storage.Test())
	testClient = &http.Client{}
	once       = sync.Once{}
)

func TestServer_Search(t *testing.T) {
	isValidURL(t, "/KR/블록체인*전자투표")
	//testFunc(t, "/search/EN/블록체인*전자투표", "[[[\"프로그램\", 0.6286733150482178], [\"컴퓨터\", 0.6282070279121399], [\"설비\", 0.625815749168396], [\"장치\", 0.5869899988174438], [\"원자로\", 0.5834010243415833], [\"체인\", 0.5802308917045593], [\"구비\", 0.5793227553367615], [\"화폐\", 0.5754896998405457], [\"투표\", 0.5733440518379211], [\"관련\", 0.5710488557815552]], [[\"원자로\", 0.6293480396270752], [\"장치\", 0.6222476959228516], [\"구비\", 0.6169931888580322], [\"체인\", 0.6099985241889954], [\"판독\", 0.6062192916870117], [\"위해\", 0.5938752293586731], [\"구현\", 0.5908535718917847], [\"관리\", 0.5875354409217834], [\"서비스\", 0.5858620405197144], [\"프로그램\", 0.585555911064148]]]")
}

func testFunc(t *testing.T, urlSuffix string, expected string) {

	once.Do(func() {
		go func() {
			err := testServer.Start()
			assert.NoError(t, err)
		}()
		time.Sleep(1 * time.Second)
	})

	response, err := testClient.Get("http://localhost:80" + urlSuffix)
	assert.NoError(t, err)

	result, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.Equal(t, expected, string(result))

}

func isValidURL(t *testing.T, urlSuffix string) {
	once.Do(func() {
		go func() {
			err := testServer.Start()
			assert.NoError(t, err)
		}()
		time.Sleep(1 * time.Second)
	})
	result, err := testClient.Get("http://localhost:80" + urlSuffix)
	assert.NoError(t, err)
	_, err = io.ReadAll(result.Body)
	assert.NoError(t, err)
}
