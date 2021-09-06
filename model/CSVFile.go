package model

import (
	"fmt"
	"os"
)

type CSVGroup struct {
	id        string
	data      []CSVUnit
	separator string
}

func NewCSV(id string) *CSVGroup {
	return &CSVGroup{id, make([]CSVUnit, 0), "\t"}
}

func (c *CSVGroup) Append(unit CSVUnit) {
	c.data = append(c.data, unit)
}

func (c *CSVGroup) File() (file *os.File, err error) {

	if file, err = os.Create(c.id + ".csv"); err != nil {
		return
	}

	_, err = fmt.Fprintln(file, "name"+"\t"+"item")

	for _, v := range c.data {
		_, err = fmt.Fprintln(file, v.Serialize(c.separator))
	}

	return

}
