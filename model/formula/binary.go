package formula

import (
	"github.com/simp7/patent-middle-server/model"
	"strings"
)

type binary struct {
	elements []model.Group
	operator string
}

func OR(elements ...model.Group) model.Binary {
	return binary{elements, "+"}
}

func AND(elements ...model.Group) model.Binary {
	return binary{elements, "*"}
}

func (b binary) Append(target model.Group) model.Binary {
	return binary{append(b.elements, target), b.operator}
}

func (b binary) String() (result string) {

	for _, v := range b.elements {
		result = result + b.operator + v.String()
	}

	result = "(" + strings.TrimPrefix(result, b.operator) + ")"

	return

}

func (b binary) First() string {
	return b.elements[0].String()
}
