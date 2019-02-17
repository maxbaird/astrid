package board

import (
	"astrid/dimension"
	"astrid/tile"
)

//Board holds all tiles
type Board struct {
	Tiles     []tile.Tile
	dimension dimension.Dimension
	Size      uint16
}

//New makes the word puzzle board from tiles
func New(tiles []tile.Tile, height uint16, width uint16) *Board {
	board := &Board{}
	board.dimension.Height = height
	board.dimension.Width = width
	board.Size = board.dimension.Height * board.dimension.Width
	board.Tiles = tiles

	var i uint16
	var j uint16

	for i = 0; i < height; i++ {
		for j = 0; j < width; j++ {
			idx := (i * width) + j
			board.Tiles[idx].Coordinate.X = j
			board.Tiles[idx].Coordinate.Y = i
			board.Tiles[idx].SetPaths(board.dimension, board.Tiles)
		}
	}
	return board
}

//PlaceLetters places letters onto the board
func (board *Board) PlaceLetters(letters string) {
	var i uint16

	for i = 0; i < board.Size; i++ {
		board.Tiles[i].Letter = rune(letters[i])
	}
}
