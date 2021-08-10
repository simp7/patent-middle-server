package main

import (
	"github.com/simp7/patent-middle-server/storage"
	"github.com/simp7/patent-middle-server/storage/cache"
	"github.com/simp7/patent-middle-server/storage/rest"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestKipris_GetClaims(t *testing.T) {

	cacheDB, err := cache.Mongo(cache.Config{URL: "mongodb://localhost", DBName: "Patent", CollectionName: "claim"})
	assert.NoError(t, err)
	restServer := rest.New(rest.Config{WordURL: "http://localhost:8080/", ClaimURL: "http://plus.kipris.or.kr/kipo-api/kipi/patUtiModInfoSearchSevice/getBibliographyDetailInfoSearch", Key: os.Getenv("KIPRIS"), Row: 500})
	server := storage.New(restServer, cacheDB)

	scenario := []struct {
		input  string
		output string
	}{
		{"3진법*반도체*누설전류*컴퓨터*!생물학", ""},
	}

	for _, v := range scenario {
		group := server.GetClaims(v.input)
		assert.NoError(t, err)
		assert.Equal(t, group, v.output)
	}

}
