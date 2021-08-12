package main

import (
	"os/exec"
	"strings"
)

type NLP struct {
	supportFile string
}

func Word2vec() NLP {
	return NLP{rootTo("word2vec.py")}
}

func LDA() NLP {
	return NLP{rootTo("LDA.py")}
}

func LSA() NLP {
	return NLP{rootTo("LSA.py")}
}

func (n NLP) Process(dataCSV string, arg ...string) ([]byte, error) {

	script := rootTo("execute.sh")
	cmd := exec.Command(script, append([]string{n.supportFile, dataCSV, "10"}, arg...)...)

	result, err := cmd.CombinedOutput()
	result = []byte(strings.TrimSpace(string(result)))

	return result, err

}
