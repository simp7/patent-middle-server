package formula

import (
	"strings"
)

type expression struct {
	result *formula
}

func Interpret(target string) *formula {

	i := &expression{New()}
	idx := 0

	blocks := strings.Split(target, "*")
	for _, block := range blocks {
		switch []rune(block)[0] {
		case '!':
			i.result.Exclude(strings.TrimPrefix(block, "!"))
		case '(':
			block = strings.Trim(block, "()")
			words := strings.Split(block, "+")
			for _, word := range words {
				i.result.Alias(idx, word)
			}
			idx++
		default:
			i.result.Alias(idx, block)
			idx++
		}
	}

	return i.result

}
