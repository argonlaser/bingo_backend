package gameDetails

type Mode int

const (
	MODE_ENDLESS = 0
	MODE_TIMER = 1
)


type Status int
const ( 
	GAME_WAIT = 0
	GAME_IN_PROGRESS = 1
	GAME_OVER = 2

	GAME_INVALID = -1 	
)
const (
	GAME_MAX_CAPACITY = 2
)