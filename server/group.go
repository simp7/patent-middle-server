package server

import (
	"strings"
)

type group struct {
	elements []group
	operator string
}

func or(elements ...group) group {
	return group{elements, "+"}
}

func and(elements ...group) group {
	return group{elements, "*"}
}

func not(target group) group {
	return group{[]group{target}, "!"}
}

func (g group) String() (result string) {

	if g.operator == "!" {
		result = "!" + g.elements[0].String()
		return
	}

	for _, v := range g.elements {
		result = result + g.operator + v.String()
	}

	result = "(" + strings.TrimPrefix(result, g.operator) + ")"

	return

}
