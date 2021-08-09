package formula

import (
	"github.com/simp7/patent-middle-server/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInterpret(t *testing.T) {
	scenario := []struct {
		desc   string
		input  string
		output *formula
	}{
		{"simple", "블록체인*전자투표", &formula{Included: []model.Binary{OR(Element("블록체인")), OR(Element("전자투표"))}, Excluded: AND()}},
		{"has alias", "블록체인*(전자투표+선거)", &formula{Included: []model.Binary{OR(Element("블록체인")), OR(Element("전자투표"), Element("선거"))}, Excluded: AND()}},
		{"has many alias", "(블록체인+알고리즘)*(전자투표+선거+다수결)", &formula{Included: []model.Binary{OR(Element("블록체인"), Element("알고리즘")), OR(Element("전자투표"), Element("선거"), Element("다수결"))}, Excluded: AND()}},
		{"has excluded", "블록체인*전자투표*!pcr", &formula{Included: []model.Binary{OR(Element("블록체인")), OR(Element("전자투표"))}, Excluded: AND(Element("pcr"))}},
		{"has many excluded", "블록체인*전자투표*!pcr*!의료*!원자력", &formula{Included: []model.Binary{OR(Element("블록체인")), OR(Element("전자투표"))}, Excluded: AND(Element("pcr"), Element("의료"), Element("원자력"))}},
		{"complicated", "(블록체인+알고리즘)*(전자투표+선거+다수결)*!pcr*!의료*!원자력", &formula{Included: []model.Binary{OR(Element("블록체인"), Element("알고리즘")), OR(Element("전자투표"), Element("선거"), Element("다수결"))}, Excluded: AND(Element("pcr"), Element("의료"), Element("원자력"))}},
	}

	for _, v := range scenario {
		assert.Equal(t, v.output, Interpret(v.input), v.desc)
	}

}
