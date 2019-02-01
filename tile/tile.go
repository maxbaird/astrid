package tile

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

//MakeTiles creates, initializes and returns n tiles
func MakeTiles(n uint16) []Tile {
	tiles := make([]Tile, n)

	var i uint16

	for i = 0; i < n; i++ {
		tiles[i].ID = i
		tiles[i].Letter = 0
		tiles[i].N = nil
		tiles[i].S = nil
		tiles[i].E = nil
		tiles[i].W = nil
		tiles[i].NE = nil
		tiles[i].SE = nil
		tiles[i].SW = nil
		tiles[i].NW = nil
		tiles[i].Coordinate.X = 0
		tiles[i].Coordinate.Y = 0
	}

	return tiles
}
