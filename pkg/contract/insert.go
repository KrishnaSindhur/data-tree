package contract

type Data struct {
	Dim []Dimensions `json:"dim"`
	Met []Metrics    `json:"metrics"`
}

type Dimensions struct {
	Key   string `json:"key"`
	Value string `json:"val"`
}

type Metrics struct {
	Key   string `json:"key"`
	Value int    `json:"val"`
}
