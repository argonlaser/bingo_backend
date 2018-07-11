package gameDetails
import(
	"errors"
)

// Player instance
type Player struct {
	Board GameBoard   `json:"board"`
	Bingoed bool      `json:"bingoed"`
	Player_id string  `json:"player_id"`
}

// NewPlayer gives new player instance
func NewPlayer(mPlayer string) *Player {
	var player = &Player{
		Player_id: mPlayer,
	}

	return player
}

func (p *Player) canBingo() bool {
	var strikedCount = 0

	for rowVal := 1; rowVal <= 5; rowVal++ {
		var isRowFullyStriked = true

		for colVal := 1; colVal <= 5; colVal++ {
			var box, err = p.Board.GetCheckBox(uint8(rowVal), uint8(colVal))
			if err == nil {
				if !box.IsStriked {
					isRowFullyStriked = false
					break
				}
			}
		}

		if isRowFullyStriked {
			strikedCount++
		}
	}

	if strikedCount >= 5 {
		return true
	}

	for colVal := 1; colVal <= 5; colVal++ {
		var isColFullyStriked = true

		for rowVal := 1; rowVal <= 5; rowVal++ {
			var box, err = p.Board.GetCheckBox(uint8(rowVal), uint8(colVal))
			if err == nil {
				if !box.IsStriked {
					isColFullyStriked = false
					break
				}
			}
		}

		if isColFullyStriked {
			strikedCount++
		}
	}

	if strikedCount >= 5 {
		return true
	}

	return false
}

// NewPlayer gives new player instance
func (p *Player) FillPlayerBoard(gb *GameBoard) {
	p.Board.FillCheckBox(gb)
}

// Bingo allows user to bingo
func (p *Player) IsBingo() error {
	if p.Bingoed {
		return errors.New("Already bingoed player. Could not bingo more than once.")
	}

	if p.canBingo() {
		p.Bingoed = true
		return nil
	}

	return errors.New("Not enough striked count")
}