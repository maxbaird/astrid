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

func traverse(tile *tile.Tile, p *path, depth int, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}

	if depth == 9 {
		return
	}

	tilePath := &path{}
	tilePath.letters = make([]rune, 9)
	tilePath.traversePath = make(map[int]struct{})

	if p != nil {
		for k, v := range p.traversePath {
			tilePath.traversePath[k] = v
		}

		copy(tilePath.letters, p.letters)
	}

	tilePath.traversePath[int(tile.ID)] = struct{}{}
	tilePath.letters[depth] = tile.Letter

	word := string(tilePath.letters[0 : depth+1])

	if lexis.IsWord(word) {
		fmt.Println(word)
	}

	if canMove(tile.N, tilePath) {
		traverse(tile.N, tilePath, depth+1, nil)
	}
	if canMove(tile.S, tilePath) {
		traverse(tile.S, tilePath, depth+1, nil)
	}
	if canMove(tile.E, tilePath) {
		traverse(tile.E, tilePath, depth+1, nil)
	}
	if canMove(tile.W, tilePath) {
		traverse(tile.W, tilePath, depth+1, nil)
	}
	if canMove(tile.NE, tilePath) {
		traverse(tile.NE, tilePath, depth+1, nil)
	}
	if canMove(tile.SE, tilePath) {
		traverse(tile.SE, tilePath, depth+1, nil)
	}
	if canMove(tile.SW, tilePath) {
		traverse(tile.SW, tilePath, depth+1, nil)
	}
	if canMove(tile.NW, tilePath) {
		traverse(tile.NW, tilePath, depth+1, nil)
	}
}

//FindWords ...
func FindWords(board *board.Board) {
	fmt.Println("Finding...")
	var wg sync.WaitGroup

	wg.Add(len(board.Tiles))
	for i := 0; i < 16; i++ {
		go traverse(&board.Tiles[i], nil, 0, &wg)
	}

	wg.Wait()
}
