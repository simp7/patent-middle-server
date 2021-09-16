package model

import (
	"time"
)

type CSVGroup struct {
	SearchWord string
	FileName   string
	Data       []CSVUnit
	Separator  string
}

func NewCSV(input string) *CSVGroup {

	fileName := time.Now().String() + "@" + input + ".csv"
	data := make([]CSVUnit, 0)

	return &CSVGroup{input, fileName, data, "\t"}

}

func (c *CSVGroup) Append(unit CSVUnit) {
	c.Data = append(c.Data, unit)
}

func (c *CSVGroup) Header() string {
	return "name" + c.Separator + "item" + "\n"
}

func (c *CSVGroup) IsEmpty() bool {
	return len(c.Data) == 0
}
