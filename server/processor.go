package server

import (
	"encoding/csv"
	"io"
)

func ProcessCSV(reader io.Reader) ([]unit, error) {

	r := csv.NewReader(reader)

	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	return convertRecords(records)

}

func convertRecords(records [][]string) (result []unit, err error) {

	result = make([]unit, len(records))

	for i := range result {
		result[i], err = convertRecord(records[i])
		if err != nil {
			return
		}
	}

	return

}

func convertRecord(record []string) (unit, error) {
	return Unit(record[0], record[1])
}
