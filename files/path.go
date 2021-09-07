package files

type Path string

const (
	ROOT         Path = ""
	CONFIG            = "config.yaml"
	INITIALIZE        = "initialize.sh"
	EXECUTE           = "execute.sh"
	WORD2VEC          = "word2vec.py"
	LDA               = "LDA.py"
	LSA               = "LSA.py"
	REQUIREMENTS      = "requirements.txt"
	LOG               = "server.log"
	STOPWORDS         = "stopwords.json"
)

func New(path string) Path {
	return Path(path)
}
