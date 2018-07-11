package gameDetails

import (
	 "fmt"
	 "errors"

	"github.com/jinzhu/copier"
)
// Contains details about the Gameboard which consists of checkboxes

type GameBoard struct {
	Checkbox [5][5]CheckBox `json:"checkbox"`
	StrikedCount uint8 `json:"striked_count"`
}

func (gb *GameBoard) IncStrikedCount() {
	gb.StrikedCount++
}

func (gb *GameBoard) Strike(val uint8) bool {
	var i,j int
	var status bool
	for i = 0; i < 5; i++ {
		for j = 0; j < 5; j++ {
			cbox := &gb.Checkbox[i][j]
			if val == cbox.Val { 
				if cbox.IsStriked == false {
					cbox.IsStriked = true
					status = true
				} else {
					fmt.Println("Already striked the checkbox")
					status = false
				}		
			} else {
				fmt.Println("What the fuck?! Wrong value sent")
				status = false
			}
		}
	}
	return status
}

func (gb *GameBoard) FillCheckBox(mGBoard *GameBoard) bool {
	copier.Copy(gb, &mGBoard)
	return true
}

func (gb *GameBoard) CountStriked() uint8 {
	return 2
}

func (gb *GameBoard) GetCheckBox(row, col uint8) (*CheckBox, error) {
		for i := 0; i < len(gb.Checkbox); i++ {
			for j := 0; j < len(gb.Checkbox[i]); j++ {
				if (gb.Checkbox[i][j].Row == row) && (gb.Checkbox[i][j].Col == col) {
					return &gb.Checkbox[i][j], nil
			}
		}
	}

	return nil, errors.New("CheckBox is not found")
}
