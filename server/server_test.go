package server

import (
	"github.com/google/logger"
	"github.com/simp7/patent-middle-server/server/claimDB"
	"net/http"
	"os"
	"testing"
)

var (
	testServer = server{http.Server{}, claimDB.Test(), logger.Init("test", true, false, os.Stdout)}
)

func TestServer_Search(t *testing.T) {
}
