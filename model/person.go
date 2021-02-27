package model

type Student struct {
	Name string   `json:"name"` // name
	Lct  Location `json:"lct"`  //location
}

//Location defines location combine x & y
type Location struct {
	X float32 `json:"x"` // x for location
	Y float32 `json:"y"` // y for location
}
