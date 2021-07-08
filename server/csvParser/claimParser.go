package csvParser

import (
	"github.com/simp7/patent-middle-server/model"
	"io"
)

type hasHeader struct {
	csvParser
}

func ParserWithHeader(reader io.Reader) hasHeader {
	return hasHeader{csvParser{reader}}
}

func (h *hasHeader) Parse() (result []model.CSVUnit, err error) {
	result, err = h.csvParser.Parse()
	if len(result) > 0 {
		result = result[1:]
	}
	return

}
