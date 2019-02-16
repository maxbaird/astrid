package blitz

import (
	"astrid/board"
	"astrid/finder"
	"astrid/tile"
	"astrid/wordcolumn"
	"fmt"
)

//Blitz ...
type Blitz struct {
	wordColumn []wordcolumn.WordColumn
	board      *board.Board
}

//New ...
func New(height uint16, width uint16) *Blitz {
	tiles := make([]tile.Tile, height*width)
	board := board.New(tiles, height, width)

	wc := make([]wordcolumn.WordColumn, board.Size)

	for i := 0; i < int(board.Size); i++ {
		wc[i].Words = make(map[string]struct{})
	}

	return &Blitz{wc, board}
}

//PrintWords ...
func (blitz *Blitz) PrintWords() {
	for _, wc := range blitz.wordColumn {
		for k := range wc.Words {
			fmt.Println(k)
		}
	}
}

//Start ...
func (blitz Blitz) Start() {
	blitz.board.PlaceLetters("abcdefghijklmnop")
	blitz.board.PrintBoard()
	finder.FindWords(blitz.board, blitz.wordColumn)
	blitz.PrintWords()
}
