package rest

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestKipris_GetClaims(t *testing.T) {

	server := New(Config{"http://plus.kipris.or.kr/kipo-api/kipi/patUtiModInfoSearchSevice/getWordSearch", "http://plus.kipris.or.kr/kipo-api/kipi/patUtiModInfoSearchSevice/getBibliographyDetailInfoSearch", os.Getenv("KIPRIS"), 500}, nil)
	scenario := []struct {
		desc   string
		number string
		name   string
	}{
		{"1", "1020130053561", "매질통신 및 이동통신 또는 인터넷을 이용한 원격자동검침 및 생활보안감시 복합서비스 제공방법 및 이를 위한 시스템"},
	}

	for _, v := range scenario {
		assert.Equal(t, v.name, server.GetClaims(v.number).Name, v.desc)
	}

}

//func TestKipris_GetNumbers(t *testing.T) {
//
//	server := New("http://plus.kipris.or.kr/kipo-api/kipi/patUtiModInfoSearchSevice/getWordSearch", "http://plus.kipris.or.kr/kipo-api/kipi/patUtiModInfoSearchSevice/getBibliographyDetailInfoSearch", os.Getenv("KIPRIS"))
//	scenario := []struct {
//		desc  string
//		input   string
//		numbers string
//	} {
//		{"1","1020130053561", ""},
//	}
//
//	for _, v := range scenario {
//		assert.Equal(t, v.numbers, <-server.GetNumbers(v.input), v.desc)
//	}
//
//}
