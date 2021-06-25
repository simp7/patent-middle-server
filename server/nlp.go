package server

import (
	"os/exec"
)

type nlp struct {
	*exec.Cmd
}

func Korean() *nlp {
	n := new(nlp)
	//n.Cmd = exec.Command("")
	return n
}

func English() *nlp {
	n := new(nlp)
	return n
}

func (n *nlp) Process(s string) {
	//n.Cmd.Output()
}
