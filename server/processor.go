package server

import (
	"encoding/csv"
	"os"
)

type processor struct {
	nlp *nlp
}

func Processor(nlpObject *nlp) *processor {
	p := new(processor)
	p.nlp = nlpObject
	return p
}

func (p *processor) getCSV() {

	csv.NewReader(os.Stdin)

}

func (p *processor) extract() {

}

func (p *processor) Process() {
	p.getCSV()
	p.extract()
}
