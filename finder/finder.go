package finder

import (
	"astrid/board"
	"astrid/lexis"
	"astrid/tile"
	"fmt"
	"sync"
)

type path struct {
	root         int
	depth        int
	letters      []rune
	traversePath map[int]struct{}
}

func canMove(tile *tile.Tile, tilePath *path) bool {
	if tile == nil {
		return false
	}

	if _, ok := tilePath.traversePath[int(tile.ID)]; ok {
		return false
	}

	return true
}

func traverse(tile *tile.Tile, p *path, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}

	if p != nil { //Will be nil on first call
		if p.depth == 9 {
			return
		}
	}

	tilePath := &path{}
	tilePath.letters = make([]rune, 9)
	tilePath.traversePath = make(map[int]struct{})

	if p == nil { //Will be nil on first call
		tilePath.root = int(tile.ID)
	} else {
		for k, v := range p.traversePath {
			tilePath.traversePath[k] = v
		}
		copy(tilePath.letters, p.letters)
		tilePath.depth = p.depth
		tilePath.root = p.root
	}

	tilePath.traversePath[int(tile.ID)] = struct{}{}
	tilePath.letters[tilePath.depth] = tile.Letter
	tilePath.depth++

	word := string(tilePath.letters[0:tilePath.depth])

	if lexis.IsWord(word) {
		fmt.Printf("%s:[%d]\n", word, tilePath.root)
	}

	if canMove(tile.N, tilePath) {
		traverse(tile.N, tilePath, nil)
	}
	if canMove(tile.S, tilePath) {
		traverse(tile.S, tilePath, nil)
	}
	if canMove(tile.E, tilePath) {
		traverse(tile.E, tilePath, nil)
	}
	if canMove(tile.W, tilePath) {
		traverse(tile.W, tilePath, nil)
	}
	if canMove(tile.NE, tilePath) {
		traverse(tile.NE, tilePath, nil)
	}
	if canMove(tile.SE, tilePath) {
		traverse(tile.SE, tilePath, nil)
	}
	if canMove(tile.SW, tilePath) {
		traverse(tile.SW, tilePath, nil)
	}
	if canMove(tile.NW, tilePath) {
		traverse(tile.NW, tilePath, nil)
	}
}

//FindWords ...
func FindWords(board *board.Board) {
	fmt.Println("Finding...")
	var wg sync.WaitGroup

	numTiles := len(board.Tiles) //TODO make this a const in board

	wg.Add(numTiles)

	for i := 0; i < numTiles; i++ {
		go traverse(&board.Tiles[i], nil, &wg)
	}

	wg.Wait()
}
