package nlp

import (
	"os/exec"
	"strings"
)

type nlp struct {
	cmd string
}

func Word2vec() nlp {
	return nlp{"./word2vec.sh"}
}

func LDA() nlp {
	return nlp{"./LDA.sh"}
}

func (n nlp) Process(tmpFile string, arg ...string) ([]byte, error) {

	cmd := exec.Command(n.cmd, append([]string{tmpFile, "10"}, arg...)...)

	result, err := cmd.CombinedOutput()
	result = []byte(strings.TrimSpace(string(result)))

	return result, err

}
