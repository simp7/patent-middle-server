package rest

type Config struct {
	WordURL  string `yaml:"word-search"`
	ClaimURL string `yaml:"claim-search"`
	Key      string `yaml:"api-key"`
	Row      int    `yaml:"row"`
}
