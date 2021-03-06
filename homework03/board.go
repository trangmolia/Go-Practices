package main

import "github.com/gdamore/tcell"

// Board is screen to run game
type Board struct {
	clist  CellList
	width  int
	height int
}

// Init sets color for cells of board game
func (b Board) Init(s tcell.Screen, cList CellList) {
	boardStyle := tcell.StyleDefault.Background(tcell.NewRGBColor(255, 255, 255))
	for i, row := range cList.grid {
		for j, cell := range row {
			cellStyle := tcell.StyleDefault.Background(tcell.NewRGBColor(0, 255, 0))
			if cell.state {
				s.SetCell(j*2, i, cellStyle, ' ')
				s.SetCell(j*2+1, i, cellStyle, ' ')
			} else {
				s.SetCell(j*2, i, boardStyle, ' ')
				s.SetCell(j*2+1, i, boardStyle, ' ')
			}
		}
	}
}
