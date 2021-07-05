package server

import (
	"encoding/csv"
	"io"
)

type processor struct {
}

func Processor() *processor {
	p := new(processor)
	return p
}

func (p *processor) Process(reader io.Reader) ([]unit, error) {

	r := csv.NewReader(reader)

	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	return p.convertRecords(records)

}

func (p *processor) convertRecords(records [][]string) (result []unit, err error) {

	result = make([]unit, len(records))

	for i := range result {
		result[i], err = p.convertRecord(records[i])
		if err != nil {
			return
		}
	}

	return

}

func (p *processor) convertRecord(record []string) (unit, error) {
	return Unit(record[0], record[1])
}
