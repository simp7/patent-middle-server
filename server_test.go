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
	//testFunc(t, "/search/KR/블록체인*전자투표", "[[0, \"0.023*\\\"센서\\\" + 0.020*\\\"신호\\\" + 0.016*\\\"컨텐츠\\\" + 0.014*\\\"장치\\\" + 0.013*\\\"무선\\\"\"], [1, \"0.026*\\\"통신\\\" + 0.022*\\\"모듈\\\" + 0.015*\\\"확인\\\" + 0.013*\\\"암호\\\" + 0.013*\\\"신원\\\"\"], [2, \"0.032*\\\"블록\\\" + 0.028*\\\"체인\\\" + 0.024*\\\"단말\\\" + 0.022*\\\"인증\\\" + 0.019*\\\"투표\\\"\"], [3, \"0.021*\\\"수신\\\" + 0.021*\\\"장치\\\" + 0.020*\\\"서버\\\" + 0.019*\\\"컴퓨터\\\" + 0.013*\\\"요청\\\"\"], [4, \"0.032*\\\"거래\\\" + 0.020*\\\"암호\\\" + 0.018*\\\"투표\\\" + 0.016*\\\"블록\\\" + 0.015*\\\"서버\\\"\"], [5, \"0.050*\\\"노드\\\" + 0.039*\\\"블록\\\" + 0.029*\\\"체인\\\" + 0.024*\\\"트랜잭션\\\" + 0.022*\\\"컴퓨터\\\"\"], [6, \"0.024*\\\"설비\\\" + 0.024*\\\"원자로\\\" + 0.023*\\\"장치\\\" + 0.022*\\\"구성\\\" + 0.021*\\\"블록\\\"\"], [7, \"0.131*\\\"단말기\\\" + 0.066*\\\"통합\\\" + 0.057*\\\"상황\\\" + 0.050*\\\"서버\\\" + 0.038*\\\"통신\\\"\"], [8, \"0.040*\\\"블록\\\" + 0.039*\\\"체인\\\" + 0.033*\\\"투표\\\" + 0.032*\\\"트랜잭션\\\" + 0.030*\\\"키\\\"\"], [9, \"0.017*\\\"결제\\\" + 0.015*\\\"간\\\" + 0.015*\\\"모바일\\\" + 0.014*\\\"액세스\\\" + 0.013*\\\"통신\\\"\"]]")
	testFunc(t, "/search/EN/블록*투표", "[[[\"프로그램\", 0.6286733150482178], [\"컴퓨터\", 0.6282070279121399], [\"설비\", 0.625815749168396], [\"장치\", 0.5869899988174438], [\"원자로\", 0.5834010243415833], [\"체인\", 0.5802308917045593], [\"구비\", 0.5793227553367615], [\"화폐\", 0.5754896998405457], [\"투표\", 0.5733440518379211], [\"관련\", 0.5710488557815552]], [[\"원자로\", 0.6293480396270752], [\"장치\", 0.6222476959228516], [\"구비\", 0.6169931888580322], [\"체인\", 0.6099985241889954], [\"판독\", 0.6062192916870117], [\"위해\", 0.5938752293586731], [\"구현\", 0.5908535718917847], [\"관리\", 0.5875354409217834], [\"서비스\", 0.5858620405197144], [\"프로그램\", 0.585555911064148]]]")
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
