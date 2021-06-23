package server

import (
	"os/exec"
)

type nlp struct {
	*exec.Cmd
}

func (n nlp) Process(s string) {
	//n.Cmd.Output()
}
