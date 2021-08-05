package model

type CSVUnit struct {
	Key   string
	Value string
}

func (c CSVUnit) Serialize() string {
	return c.Key + ",\"" + c.Value + "\""
}
