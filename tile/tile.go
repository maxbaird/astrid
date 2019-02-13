package tile

import (
	"astrid/dimension"
	"fmt"
)

type coordinate struct {
	X uint16
	Y uint16
}

//Tile represents a letter on the board
type Tile struct {
	ID         uint16
	Letter     rune
	Coordinate coordinate
	N          *Tile
	S          *Tile
	E          *Tile
	W          *Tile
	NE         *Tile
	SE         *Tile
	SW         *Tile
	NW         *Tile
}

//SetPaths ...
func (tile *Tile) SetPaths(d dimension.Dimension, t []Tile) {
	height := int(d.Height)
	width := int(d.Width)

	X := int(tile.Coordinate.X)
	Y := int(tile.Coordinate.Y)

	idx := (height * Y) + X //Index of tile
	tile.ID = uint16(idx) + 1

	NIdx := idx - width
	SIdx := idx + height
	EIdx := idx + 1
	WIdx := idx - 1
	NEIdx := idx - (width - 1)
	SEIdx := idx + (width + 1)
	SWIdx := idx + (width - 1)
	NWIdx := idx - (width + 1)

	if !(NIdx < 0) {
		tile.N = &t[NIdx]
	}

	if !(Y+1 >= height) {
		tile.S = &t[SIdx]
	}

	if !(X+1 >= width) {
		tile.E = &t[EIdx]
	}

	if !(X-1 < 0) {
		tile.W = &t[WIdx]
	}

	if !(NIdx < 0 || X+1 >= width) {
		tile.NE = &t[NEIdx]
	}

	if !(Y+1 >= height || X+1 >= width) {
		tile.SE = &t[SEIdx]
	}

	if !(Y+1 >= height || X-1 < 0) {
		tile.SW = &t[SWIdx]
	}

	if !(NIdx < 0 || X-1 < 0) {
		tile.NW = &t[NWIdx]
	}
}

//PrintTile ...
func (tile *Tile) PrintTile() {
	fmt.Printf("[%c](%d,%d)", tile.Letter, tile.Coordinate.X, tile.Coordinate.Y)
	if tile.N != nil {
		fmt.Printf(" N:%c", tile.N.Letter)
	}
	if tile.S != nil {
		fmt.Printf(" S:%c", tile.S.Letter)
	}
	if tile.E != nil {
		fmt.Printf(" E:%c", tile.E.Letter)
	}
	if tile.W != nil {
		fmt.Printf(" W:%c", tile.W.Letter)
	}
	if tile.NE != nil {
		fmt.Printf(" NE:%c", tile.NE.Letter)
	}
	if tile.SE != nil {
		fmt.Printf(" SE:%c", tile.SE.Letter)
	}
	if tile.SW != nil {
		fmt.Printf(" SW:%c", tile.SW.Letter)
	}
	if tile.NW != nil {
		fmt.Printf(" NW:%c", tile.NW.Letter)
	}
}
