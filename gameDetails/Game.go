package gameDetails

import (
	"github.com/satori/go.uuid"

	"time"
	"fmt"
	"errors"
)



// Game instance
type Game struct {
	Game_id            string         `json:"game_id"`
	Creator            *Player        `json:"creator"`
	Players            []*Player      `json:"players"`
	GameLoop           *time.Timer    `json:"gameloop"`
	CurrentPlayer      *Player        `json:"current_player"`
	currentPlayerIndex int            `json:"current_player_index"`
	Results            []*Result      `json:"results"`
	Status				Status        `json:"status"`
	Created_at         time.Time      `json:"created_at"`
	Mode               Mode           `json:"mode"`
}

// Create newgame
func NewGame() *Game {
	// error handling
	u4, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		panic(err)
	}
	var g = &Game{
		Game_id: u4.String(),
	}
	return g
}

// Start the game
func (g *Game) Start() {
	g.currentPlayerIndex = -1
	g.startNextTurn()
}

// startNextTurn will give the turn to next user on line
func (g *Game) startNextTurn() {
	var playerIndex = g.currentPlayerIndex + 1
	if playerIndex == len(g.Players) {
		playerIndex = 0
	}

	g.currentPlayerIndex = playerIndex
	g.CurrentPlayer = g.Players[playerIndex]

}

// PlayerStrike used to strike a box in the player's board
func (g *Game) PlayerStrike(p *Player, val uint8) error {
	if p != g.CurrentPlayer {
		return errors.New("Could not strike, as its not the turn of the player")
	}

	// strike the number in player's board
	status := p.Board.Strike(val)

	if status != false {
		return nil
	}

	// start the next turn
	g.startNextTurn()
	return nil
}

// // PlayerBingo executes for a player, when he tries to press bingo in frontend
// func (g *Game) PlayerBingo(p *Player) (*Result, error) {
// 	err := p.Bingo()

// 	if err != nil {
// 		return nil, err
// 	}

// 	var result = &Result{
// 		Position: int32(len(g.Results)),
// 		Player:   p,
// 	}

// 	g.Results = append(g.Results, result)

// 	return result, nil
// }

// GetPlayerByID gets player in a game by ID
func (g *Game) GetPlayerByID(playerID string) (*Player, error) {
	var reqPlayer *Player

	for _, p := range g.Players {
		if p.Player_id == playerID {
			reqPlayer = p
			break
		}
	}

	if reqPlayer == nil {
		return nil, errors.New("PlayerNotFound")
	}
	return reqPlayer, nil
}

