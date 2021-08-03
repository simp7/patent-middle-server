package formula

import (
	"strings"
)

type expression struct {
	wordCache  []rune
	result     *formula
	input      []rune
	stringIdx  int
	includeIdx int
}

func Interpret(target string) *formula {

	i := &expression{make([]rune, 0), New(), []rune(target), 0, 0}
	strLen := len(i.input)
	idx := 0

	for idx < strLen && i.input[idx] != '-' {
		char := i.input[idx]
		idx++
		switch char {
		case '(':
			continue
		case ')':
			continue
		case '+':
			i.result.Alias(i.includeIdx, string(i.wordCache))
			i.wordCache = make([]rune, 0)
		case '*':
			i.result.Alias(i.includeIdx, string(i.wordCache))
			i.wordCache = make([]rune, 0)
			i.includeIdx++
		default:
			i.wordCache = append(i.wordCache, char)
		}
	}

	i.result.Alias(i.includeIdx, string(i.wordCache))

	if idx != strLen {

		rest := string(i.input[idx+1:])
		excluded := strings.Split(rest, "-")

		for _, word := range excluded {
			i.result.Exclude(word)
		}

	}

	return i.result

}
