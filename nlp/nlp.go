package nlp

import (
	"os/exec"
)

type nlp struct {
	cmd string
}

func Word2vec() nlp {
	//TODO: 국내 특허 자연어 처리 명령어 삽입
	return nlp{"python word2vec.py"}
}

func LDP() nlp {
	//TODO: 해외 특허 자연어 처리 명령어 삽입
	return nlp{"python ldp.py"}
}

func (n nlp) Process(tmpFile string) (string, error) {
	result, err := exec.Command(n.cmd, tmpFile).Output()
	return string(result), err
}
