package board

import (
	"blitzkrieg/tile"
)

type dimension struct {
	height uint16
	width  uint16
}

//Board holds all tiles
type Board struct {
	tiles     []tile.Tile
	dimension dimension
}

//MakeBoard makes the word puzzle board from tiles
func MakeBoard(tiles []tile.Tile, height uint16, width uint16) Board {
	tiles[0].Letter = 'c'
	var b Board
	return b
}
