package blitz

import (
	"astrid/board"
	"astrid/tile"
)

const height uint16 = 4
const width uint16 = 4

//Blitz ...
func Blitz() {
	board := new(board.Board)
	tiles := make([]tile.Tile, height*width)

	board.MakeBoard(tiles, height, width)
	board.PlaceLetters("abcdefghijklmnop")
	board.PrintBoard()
}
