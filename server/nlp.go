package server

import (
	"fmt"
	"os/exec"
	"strings"
)

type nlp struct {
	*exec.Cmd
	processor *processor
}

func Korean() *nlp {
	n := new(nlp)
	n.processor = Processor()
	//n.Cmd = exec.Command("")
	return n
}

func English() *nlp {
	n := new(nlp)
	n.processor = Processor()
	//n.Cmd = exec.Command("")
	return n
}

func (n *nlp) Process(s string) (string, error) {
	//n.Cmd.Output()
	u, err := n.processor.Process(strings.NewReader(s))
	fmt.Println(u)
	//TODO: 유저 입력을 받아서 새로운 검색식 만들기
	return s, err
}
