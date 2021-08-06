package claimDB

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestKipris_GetClaims(t *testing.T) {

	x, err := New("http://localhost:8080/", "http://plus.kipris.or.kr/kipo-api/kipi/patUtiModInfoSearchSevice/getBibliographyDetailInfoSearch", os.Getenv("KIPRIS"))
	assert.NoError(t, err)

	scenario := []struct {
		input  string
		output string
	}{
		{"3진법*반도체*누설전류*컴퓨터*!생물학", ""},
	}

	for _, v := range scenario {
		a, err := x.GetClaims(v.input)
		assert.NoError(t, err)
		assert.Equal(t, a, v.output)
	}

}
