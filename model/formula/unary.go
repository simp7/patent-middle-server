package formula

import (
	"github.com/simp7/patent-middle-server/model"
)

type unary struct {
	element  model.Group
	operator string
}

func NOT(target model.Group) unary {
	return unary{target, "!"}
}

func (u unary) String() string {
	return "!" + u.element.String()
}
