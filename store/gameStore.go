package store

import "../gameDetails"

// GameStore will contain the reference to the games
// or methods to reach those game references
type GameStore interface {
	Add(g *gameDetails.Game) error
	GetByGameID(gameID string) (*gameDetails.Game, error)
	Remove(g *gameDetails.Game) error
	Allocate(player *gameDetails.Player) (*gameDetails.Game, error)
}