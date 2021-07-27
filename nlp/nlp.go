package nlp

import (
	"fmt"
	"os/exec"
)

type nlp struct {
	cmd string
}

func Word2vec() nlp {
	//TODO: 국내 특허 자연어 처리 명령어 삽입
	return nlp{"./word2vec.sh"}
}

func LDA() nlp {
	//TODO: 해외 특허 자연어 처리 명령어 삽입
	return nlp{"./LSA.sh"}
}

func (n nlp) Process(tmpFile string) ([]byte, error) {

	cmd := exec.Command(n.cmd, tmpFile, "10")

	result, err := cmd.CombinedOutput()
	fmt.Println(string(result))
	return result, err
}
