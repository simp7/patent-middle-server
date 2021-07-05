package formula

import (
	"github.com/simp7/patent-middle-server/model"
)

type element string

func Element(s string) element {
	return element(s)
}

func (e element) String() string {
	return string(e)
}

func (e element) Append(_ model.Group) model.Group {
	return e
}
