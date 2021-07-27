package model

type Binary interface {
	Group
	Append(Group) Group
}

type Group interface {
	String() string
}

type Formula interface {
	String() string
	Exclude(string)
	Alias(int, string)
	Verify() error
}
