package main

import (
	claimDB2 "github.com/simp7/patent-middle-server/claimDB"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"sync"
	"testing"
	"time"
)

var (
	testServer = New(80, claimDB2.Test())
	testClient = &http.Client{}
	once       = sync.Once{}
)

func TestServer_Search(t *testing.T) {
	testFunc(t, "/search?formula=블록체인*전자투표&country=KR", "블록체인*전자투표")
}

func TestServer_Welcome(t *testing.T) {
	testFunc(t, "", "<h1>Hello, world!</h1>")
}

func testFunc(t *testing.T, urlSuffix string, expected string) {

	once.Do(func() {
		go testServer.Start()
		time.Sleep(1 * time.Second)
	})

	response, err := testClient.Get("http://localhost:80" + urlSuffix)
	assert.NoError(t, err)

	result, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.Equal(t, expected, string(result))

}
