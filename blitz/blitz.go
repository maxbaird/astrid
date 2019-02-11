package blitz

import (
	"astrid/board"
	"astrid/tile"
)

//Blitz ...
func Blitz() {
	board := board.Board{}
	tiles := tile.MakeTiles(16)

	board.MakeBoard(tiles, 4, 4)
	board.PlaceLetters("abcdefghijklmnop")
	board.PrintBoard()
}
