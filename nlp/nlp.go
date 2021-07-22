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
	return nlp{"nlp/LSA.py"}
}

func LDA() nlp {
	//TODO: 해외 특허 자연어 처리 명령어 삽입
	return nlp{"nlp/LSA.py"}
}

func (n nlp) Process(tmpFile string) ([]byte, error) {

	venv := exec.Command("zsh", "-c", "source "+"./venv/bin/activate")

	err := venv.Run()
	fmt.Println(err)
	cmd := exec.Command("python3", n.cmd, tmpFile, "10")

	result, err := cmd.CombinedOutput()
	fmt.Println(string(result))
	return result, err
}
