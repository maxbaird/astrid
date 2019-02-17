package blitz

import (
	"astrid/board"
	"astrid/finder"
	"astrid/printer"
	"astrid/tile"
	"astrid/wordcolumn"
)

//Blitz ...
type Blitz struct {
	WordColumn []wordcolumn.WordColumn
	Board      *board.Board
}

//New ...
func New(height int, width int) *Blitz {
	tiles := make([]tile.Tile, height*width)
	board := board.New(tiles, height, width)

	wc := make([]wordcolumn.WordColumn, board.Size)

	for i := 0; i < board.Size; i++ {
		wc[i].Words = make(map[string]struct{})
	}

	return &Blitz{wc, board}
}

//Start ...
func (blitz Blitz) Start() {
	blitz.Board.PlaceLetters("abcdefghijklmnop")
	finder.FindWords(blitz.Board, blitz.WordColumn)
	printer.PrintWords(blitz.Board, blitz.WordColumn)
}
