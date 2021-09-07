package model

type CSVGroup struct {
	ID        string
	Data      []CSVUnit
	Separator string
}

func NewCSV(id string) *CSVGroup {
	return &CSVGroup{id, make([]CSVUnit, 0), "\t"}
}

func (c *CSVGroup) Append(unit CSVUnit) {
	c.Data = append(c.Data, unit)
}
