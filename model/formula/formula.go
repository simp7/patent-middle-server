package formula

import (
	"errors"
	"github.com/simp7/patent-middle-server/model"
	"strings"
)

var shouldIncludeWord = errors.New("formula should include at least one word")

type formula struct {
	Included []model.Binary
	Excluded model.Binary
}

func Interpret(target string) model.Formula {

	i := newFormula()
	idx := 0

	blocks := strings.Split(target, "*")
	for _, block := range blocks {
		switch []rune(block)[0] {
		case '!':
			i.Exclude(strings.TrimPrefix(block, "!"))
		case '(':
			block = strings.Trim(block, "()")
			words := strings.Split(block, "+")
			for _, word := range words {
				i.Alias(idx, word)
			}
			idx++
		default:
			i.Alias(idx, block)
			idx++
		}
	}

	return i

}

func newFormula() *formula {

	f := new(formula)
	f.Included = make([]model.Binary, 0)
	f.Excluded = AND()

	return f

}

func binaryToGroup(binary []model.Binary) (group []model.Group) {
	group = make([]model.Group, len(binary))
	for i, v := range binary {
		group[i] = v
	}
	return
}

func (f *formula) String() string {

	result := AND(binaryToGroup(f.Included)...)

	if f.Excluded != nil {
		result = AND(result, NOT(OR(f.Excluded)))
	}

	return result.String()

}

func (f *formula) Exclude(target string) {
	f.Excluded = f.Excluded.Append(Element(target))
}

func (f *formula) Alias(idx int, target string) {
	if len(f.Included) <= idx {
		f.Included = append(f.Included, OR(Element(target)))
		return
	}
	f.Included[idx] = f.Included[idx].Append(Element(target))
}

func (f *formula) Verify() error {
	if len(f.Included) == 0 {
		return shouldIncludeWord
	}
	return nil
}

func (f *formula) Keywords() []string {
	result := make([]string, len(f.Included))
	for i := range f.Included {
		result[i] = f.Included[i].First()
	}
	return result
}
