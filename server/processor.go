package server

import (
	"encoding/csv"
	"io"
)

type processor struct {
	reader io.Reader
}

func Processor(reader io.Reader) *processor {
	p := new(processor)
	p.reader = reader
	return p
}

func (p *processor) Process() ([]unit, error) {
	return p.extract(p.reader)
}

func (p *processor) extract(reader io.Reader) ([]unit, error) {

	r := csv.NewReader(reader)
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	return p.convertRecords(records)

}

func (p *processor) convertRecords(records [][]string) (result []unit, err error) {

	result = make([]unit, len(records)-1)

	for i := range result {
		result[i], err = p.convertRecord(records[i+1])
		if err != nil {
			return
		}
	}

	return

}

func (p *processor) convertRecord(record []string) (unit, error) {
	return Unit(record[0], record[1])
}
