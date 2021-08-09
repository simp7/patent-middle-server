package model

import (
	"fmt"
	"os"
)

type CSVGroup struct {
	id   string
	data []CSVUnit
}

func NewCSV(id string) *CSVGroup {
	return &CSVGroup{id, make([]CSVUnit, 0)}
}

func (c *CSVGroup) Append(unit CSVUnit) {
	c.data = append(c.data, unit)
}

func (c *CSVGroup) File() (file *os.File, err error) {

	file, err = os.Create(c.id + ".csv")
	if err != nil {
		return
	}

	for _, v := range c.data {
		_, err := fmt.Fprintln(file, v.Serialize())
		if err != nil {
			continue
		}
	}

	return

}
