package server

import (
	"fmt"
	"os"
	"os/exec"
)

type nlp struct {
	*exec.Cmd
	processor *processor
}

func Korean() *nlp {
	n := new(nlp)
	n.processor = Processor(os.Stdin)
	//n.Cmd = exec.Command("")
	return n
}

func English() *nlp {
	n := new(nlp)
	n.processor = Processor(os.Stdin)
	return n
}

func (n *nlp) Process(s string) string {
	//n.Cmd.Output()
	u, _ := n.processor.Process()
	fmt.Println(u)
	//TODO: 유저 입력을 받아서 새로운 검색식 만들기
	return s
}
