package cache

type Config struct {
	URL            string `yaml:"address"`
	DBName         string `yaml:"database"`
	CollectionName string `yaml:"collection"`
}
