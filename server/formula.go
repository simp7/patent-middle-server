package server

type Formula struct {
	Included []group
	Excluded group
}

func (f Formula) String() string {
	return and(and(f.Included...), not(or(f.Excluded))).String()
}
