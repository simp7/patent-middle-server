package csvParser

import (
	"encoding/csv"
	"github.com/simp7/patent-middle-server/model"
	"io"
)

type csvParser struct {
	io.Reader
}

func Parser(reader io.Reader) csvParser {
	return csvParser{reader}
}

func (cp csvParser) Records() (records [][]string, err error) {

	reader := csv.NewReader(cp)
	records, err = reader.ReadAll()

	return

}

func (cp csvParser) Parse() ([]model.CSVUnit, error) {

	records, err := cp.Records()
	if err != nil {
		return nil, err
	}

	return convertRecords(records), err

}

func convertRecords(records [][]string) (result []model.CSVUnit) {

	result = make([]model.CSVUnit, len(records))

	for i := range result {
		result[i] = convertRecord(records[i])
	}

	return

}

func convertRecord(record []string) model.CSVUnit {
	return model.CSVUnit{Key: record[0], Value: record[1]}
}
