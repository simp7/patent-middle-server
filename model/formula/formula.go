package formula

import (
	"errors"
	"github.com/simp7/patent-middle-server/model"
)

var shouldIncludeWord = errors.New("formula should include at least one word")

type formula struct {
	Included []model.Binary
	Excluded model.Binary
}

func New() *formula {

	f := new(formula)
	f.Included = make([]model.Binary, 0)

	return f
}

func Interpret(target string) *formula {

	f := New()

	return f

}

func binaryToGroup(binary []model.Binary) (group []model.Group) {
	group = make([]model.Group, len(binary))
	for i, v := range binary {
		group[i] = v
	}
	return
}

func (f formula) String() string {

	result := AND(binaryToGroup(f.Included)...)

	if f.Excluded != nil {
		result = AND(result, NOT(OR(f.Excluded)))
	}

	return result.String()

}

func (f formula) Exclude(target string) {
	f.Excluded.Append(Element(target))
}

func (f formula) Alias(idx int, target string) {
	f.Included[idx].Append(Element(target))
}

func (f formula) Verify() error {
	if len(f.Included) == 0 {
		return shouldIncludeWord
	}
	return nil
}
