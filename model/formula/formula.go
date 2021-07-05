package formula

import (
	"github.com/simp7/patent-middle-server/model"
)

type Formula struct {
	Included []model.Binary
	Excluded model.Binary
}

func binaryToGroup(binary []model.Binary) (group []model.Group) {
	group = make([]model.Group, len(binary))
	for i, v := range binary {
		group[i] = v
	}
	return
}

func (f Formula) String() string {
	return AND(AND(binaryToGroup(f.Included)...), NOT(OR(f.Excluded))).String()
}

func (f Formula) Exclude(target string) {
	f.Excluded.Append(Element(target))
}

func (f Formula) Alias(idx int, target string) {
	f.Included[idx].Append(Element(target))
}
