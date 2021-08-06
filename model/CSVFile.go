package model

import (
	"fmt"
	"os"
)

type CSVGroup struct {
	id   string
	data []CSVUnit
}

func NewCSV(id string, data []CSVUnit) CSVGroup {
	return CSVGroup{id, data}
}

func (c CSVGroup) File() (file *os.File, err error) {

	file, err = os.Create(c.id + ".csv")
	if err != nil {
		for _, v := range c.data {
			fmt.Fprintln(file, v.Serialize())
		}
	}

	return

}
