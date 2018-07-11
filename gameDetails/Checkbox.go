package gameDetails

// CheckBox instance
// Access the member variables only through exposed methods
type CheckBox struct {
	IsStriked bool `json:"isStriked"`
	Row uint8 `json:"row"`
	Col uint8  `json:"col"`
	Val uint8 `json:"val"`
}
