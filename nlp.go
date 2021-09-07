package main

import (
	"os/exec"
	"strings"
)

type NLP func(args ...string) *exec.Cmd

func (n NLP) Process(dataCSV string, arg ...string) ([]byte, error) {

	cmd := n(append([]string{dataCSV, "10"}, arg...)...)

	result, err := cmd.CombinedOutput()
	result = []byte(strings.TrimSpace(string(result)))

	return result, err

}
