package contract

type Tree struct {
	Level1 Node1
	Level2 []Node2
	Level3 []Node3
}

type Node1 struct {
	WebReq int
	TimeSpent int
}

type Node2 struct {
	Country string
	Node1
}

type Node3 struct {
	Device string
	Node2
}

