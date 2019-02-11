package blitz

import (
	"blitzkrieg/board"
	"blitzkrieg/tile"
)

//Blitz ...
func Blitz() {
	board := board.Board{}
	tiles := tile.MakeTiles(16)

	board.MakeBoard(tiles, 4, 4)
	board.PlaceLetters("abcdefghijklmnop")
	board.PrintBoard()
}
