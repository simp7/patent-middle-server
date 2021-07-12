package nlp

import (
	"encoding/csv"
	"os/exec"
)

type nlp struct {
	cmd string
}

func Korean() nlp {
	//TODO: 국내 특허 자연어 처리 명령어 삽입
	return nlp{""}
}

func English() nlp {
	//TODO: 해외 특허 자연어 처리 명령어 삽입
	return nlp{""}
}

func (n nlp) Process(s string) (string, error) {
	result, err := exec.Command(n.cmd, s).Output()
	csv.NewWriter(exec.Command(n.cmd, s).Stdout)
	return string(result), err
}
