package blitz

import (
	"astrid/board"
	"astrid/finder"
	"astrid/tile"
)

type wordColumn struct {
	tileIndex      uint16
	wordCount      uint16
	longestWordLen uint16
	words          []string
}

//Blitz ...
type Blitz struct {
	wc    []wordColumn
	board *board.Board
}

//New ...
func New(height uint16, width uint16) *Blitz {
	tiles := make([]tile.Tile, height*width)
	board := board.New(tiles, height, width)

	wc := make([]wordColumn, board.Size)
	return &Blitz{wc, board}
}

//AddWord ...
func AddWord(word string, root uint16) {

}

//Start ...
func (blitz Blitz) Start() {
	blitz.board.PlaceLetters("abcdefghijklmnop")
	blitz.board.PrintBoard()
	finder.FindWords(blitz.board)
}
