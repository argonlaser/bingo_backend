package gameDetails

// Result denotes a game result object
type Result struct {
	Rank int8    `json:"rank"`
	Player   *Player `json:"player"`
}