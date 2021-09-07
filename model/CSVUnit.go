package model

type CSVUnit struct {
	Key   string
	Value string
}

func (c CSVUnit) Serialize(sep string) string {
	return c.Key + sep + "\"" + c.Value + "\"" + "\n"
}
