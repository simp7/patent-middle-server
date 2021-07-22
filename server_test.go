package main

import (
	"github.com/simp7/patent-middle-server/claimDB"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"sync"
	"testing"
	"time"
)

var (
	testServer = New(80, claimDB.Test())
	testClient = &http.Client{}
	once       = sync.Once{}
)

func TestServer_Search(t *testing.T) {
	testFunc(t, "/search/KR/블록체인*전자투표", "[[\"상황도\", \"이상\", \"어느\", \"통합정보를\", \"서버와\", \"통신정보\", \"단말기간\", \"단말기와\", \"포함된\", \"통신정보를\"], [\"모바일단말기\", \"블록체인\", \"특징으로\", \"것을\", \"멤버쉽정보\", \"본문결제정보\", \"정보\", \"사용자\", \"데이터\", \"복수의\"], [\"모바일단말기\", \"멤버쉽정보\", \"본문결제정보\", \"하나를\", \"정보\", \"식별정보를\", \"결제정보\", \"모바일\", \"요청된\", \"수행하는\"], [\"연계된\", \"식별정보\", \"하느님\", \"단말기\", \"중계단말기\", \"클라우드컴퓨팅\", \"스크린\", \"수신하여\", \"전송하는\", \"클라우드\"], [\"암호\", \"거래\", \"화폐의\", \"등급\", \"서비스\", \"기반\", \"제공\", \"등급의\", \"오너\", \"블록체인\"], [\"암호\", \"거래\", \"화폐의\", \"컴퓨터\", \"등급\", \"트랜잭션\", \"블록\", \"컴퓨팅\", \"등급의\", \"구현\"], [\"삭제\", \"데이터\", \"암호\", \"사용자\", \"컨텐츠\", \"획득하는\", \"화폐의\", \"탐색값을\", \"장치\", \"센서\"], [\"삭제\", \"투표\", \"특징으로\", \"것을\", \"전자\", \"유권자\", \"투표자\", \"관리\", \"전자투표\", \"인증\"], [\"데이터\", \"투표\", \"획득하는\", \"탐색값을\", \"사용자\", \"분할된\", \"인증\", \"데이터를\", \"저장된\", \"인스트럭션을\"], [\"컨텐츠\", \"시스템\", \"통신\", \"전자\", \"서비스\", \"제어\", \"질의\", \"제공\", \"가상\", \"미디어\"]]\n")
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
