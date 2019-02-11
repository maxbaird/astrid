package board

import (
	"blitzkrieg/tile"
	"fmt"
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

func (board *Board) setPaths(t *tile.Tile) {
	height := int(board.dimension.height)
	width := int(board.dimension.width)

	X := int(t.Coordinate.X)
	Y := int(t.Coordinate.Y)

	idx := (height * Y) + X //Index of tile

	NIdx := idx - width
	SIdx := idx + height
	EIdx := idx + 1
	WIdx := idx - 1
	NEIdx := idx - (width - 1)
	SEIdx := idx + (width + 1)
	SWIdx := idx + (width - 1)
	NWIdx := idx - (width + 1)

	if NIdx < 0 {
		t.N = nil
	} else {
		t.N = &board.tiles[NIdx]
	}

	if Y+1 >= height {
		t.S = nil
	} else {
		t.S = &board.tiles[SIdx]
	}

	if X+1 >= width {
		t.E = nil
	} else {
		t.E = &board.tiles[EIdx]
	}

	if X-1 < 0 {
		t.W = nil
	} else {
		t.W = &board.tiles[WIdx]
	}

	if NIdx < 0 || X+1 >= width {
		t.NE = nil
	} else {
		t.NE = &board.tiles[NEIdx]
	}

	if Y+1 >= height || X+1 >= width {
		t.SE = nil
	} else {
		t.SE = &board.tiles[SEIdx]
	}

	if Y+1 >= height || X-1 < 0 {
		t.SW = nil
	} else {
		t.SW = &board.tiles[SWIdx]
	}

	if NIdx < 0 || X-1 < 0 {
		t.NW = nil
	} else {
		t.NW = &board.tiles[NWIdx]
	}
}

//MakeBoard makes the word puzzle board from tiles
func (board *Board) MakeBoard(tiles []tile.Tile, height uint16, width uint16) {
	board.dimension.height = height
	board.dimension.width = width
	board.tiles = tiles

	var i uint16
	var j uint16

	for i = 0; i < height; i++ {
		for j = 0; j < width; j++ {
			board.tiles[j].Coordinate.X = j
			board.tiles[j].Coordinate.Y = i
			board.setPaths(&board.tiles[j])
		}
	}
}

//PlaceLetters places letters onto the board
func (board *Board) PlaceLetters(letters string) {
	var i uint16

	for i = 0; i < board.GetBoardSize(); i++ {
		board.tiles[i].Letter = rune(letters[i])
	}
}

//GetBoardSize returns the size of the board
func (board Board) GetBoardSize() uint16 {
	return board.dimension.height * board.dimension.width
}

//PrintBoard ...
func (board *Board) PrintBoard() {
	fmt.Print("Printing board\n")
	var i uint16

	for i = 0; i < board.GetBoardSize(); i++ {
		fmt.Printf("%c\n", board.tiles[i].Letter)
	}
}
