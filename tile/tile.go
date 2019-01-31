package tile

type coordinate struct {
	x uint16
	y uint16
}

//Tile represents a letter on the board
type Tile struct {
	id     uint16
	letter rune
	coord  coordinate
	N      *Tile
	S      *Tile
	E      *Tile
	W      *Tile
	NE     *Tile
	SE     *Tile
	SW     *Tile
	NW     *Tile
}

//MakeTiles creates, initializes and returns n tiles
func MakeTiles(n uint16) []Tile {
	tiles := make([]Tile, n)

	var i uint16

	for i = 0; i < n; i++ {
		tiles[i].id = i
		tiles[i].letter = 0
		tiles[i].N = nil
		tiles[i].S = nil
		tiles[i].E = nil
		tiles[i].W = nil
		tiles[i].NE = nil
		tiles[i].SE = nil
		tiles[i].SW = nil
		tiles[i].NW = nil
	}

	return tiles
}
