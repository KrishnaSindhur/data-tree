package contract

type Node struct {
	MetaData  string
	WebReq    int
	TimeSpent int
	Children  []*Node
}
