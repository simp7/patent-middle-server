package csvParser

import (
	"github.com/simp7/patent-middle-server/model"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCsvParser_Parse(t *testing.T) {

	scenario := []struct {
		testFile string
		output   []model.CSVUnit
		desc     string
	}{
		{
			"hello_world.csv",
			[]model.CSVUnit{
				{"programming", "0.999997"},
				{"terminal", "0.789901"},
				{"computer", "0.700111"},
				{"matrix", "0.450001"},
				{"C", "0.288889"},
				{"defect", "0.213841"}},
			"normal csv",
		},
	}

	for _, v := range scenario {

		var result []model.CSVUnit

		reader, err := os.Open(v.testFile)
		if err == nil {
			result, err = Parser(reader).Parse()
		}

		assert.NoError(t, err)
		assert.Equal(t, v.output, result, v.desc)

	}

}
