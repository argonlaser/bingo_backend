package store

import (
	"errors"

	"../gameDetails"
)

// LocalGameStore stores the game data in local hard drive
// if the go process exits, then all the game data will be gone
type LocalGameStore struct {
	gameList []*gameDetails.Game
}

// Add adds the game to the store
func (s *LocalGameStore) Add(g *gameDetails.Game) error {
	s.gameList = append(s.gameList, g)
	return nil
}

// GetByGameID gets game by gameID
func (s *LocalGameStore) GetByGameID(gameID string) (*gameDetails.Game, error) {
	var reqGame *gameDetails.Game

	for _, g := range s.gameList {
		if g.Game_id == gameID {
			reqGame = g
			break
		}
	}

	if reqGame == nil {
		return nil, errors.New("GameNotFound")
	}

	return reqGame, nil
}

// Remove removes the game from the store
func (s *LocalGameStore) Remove(g *gameDetails.Game) error {
	s.gameList = append(s.gameList, g)
	return nil
}

// Allocate game for the player
func (s *LocalGameStore) Allocate(p *gameDetails.Player) (*gameDetails.Game, error) {
	var g *gameDetails.Game

	for _, game:= range s.gameList {
		if (len(game.Players) < gameDetails.GAME_MAX_CAPACITY ) {
			g = game
			game.Players = append(game.Players, p)			
			return g, nil
		}
	}

	g = gameDetails.NewGame()
	s.gameList = append(s.gameList, g)
	return g, nil
}