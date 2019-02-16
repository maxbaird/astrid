package blitz

import (
	"astrid/board"
	"astrid/finder"
	"astrid/tile"
	"astrid/wordcolumn"
)

//Blitz ...
type Blitz struct {
	wc    []wordcolumn.WordColumn
	board *board.Board
}

//New ...
func New(height uint16, width uint16) *Blitz {
	tiles := make([]tile.Tile, height*width)
	board := board.New(tiles, height, width)

	wc := make([]wordcolumn.WordColumn, board.Size)
	return &Blitz{wc, board}
}

//Start ...
func (blitz Blitz) Start() {
	blitz.board.PlaceLetters("abcdefghijklmnop")
	blitz.board.PrintBoard()
	finder.FindWords(blitz.board)
}
