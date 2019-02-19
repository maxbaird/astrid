package finder

import (
	"github.com/maxbaird/astrid/board"
	"github.com/maxbaird/astrid/configuration"
	"github.com/maxbaird/astrid/lexis"
	"github.com/maxbaird/astrid/tile"
	"github.com/maxbaird/astrid/wordcolumn"
	"sync"
)

var wordColumn []wordcolumn.WordColumn

type path struct {
	root         int
	depth        int
	letters      []rune
	traversePath map[int]struct{}
}

func canMove(tile *tile.Tile, tp *path) bool {
	if tile == nil {
		return false
	}

	if _, ok := tp.traversePath[tile.ID]; ok {
		return false
	}

	return true
}

func traverse(tile *tile.Tile, p *path, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}

	if p != nil { //Will be nil on first call
		if p.depth == configuration.Config.MaxWordLength {
			return
		}
	}

	tp := &path{}
	tp.letters = make([]rune, configuration.Config.MaxWordLength)
	tp.traversePath = make(map[int]struct{})

	if p == nil { //Will be nil on first call
		tp.root = tile.ID
	} else {
		for k, v := range p.traversePath {
			tp.traversePath[k] = v
		}
		copy(tp.letters, p.letters)
		tp.depth = p.depth
		tp.root = p.root
	}

	// tp.traversePath is a map that does not hold any values
	// because we are only interested in unique keys
	tp.traversePath[tile.ID] = struct{}{}
	tp.letters[tp.depth] = tile.Letter
	tp.depth++

	word := string(tp.letters[0:tp.depth])

	if lexis.IsWord(word) {
		wordColumn[tp.root].AddWord(word, tp.root)
	}

	if canMove(tile.N, tp) {
		traverse(tile.N, tp, nil)
	}
	if canMove(tile.S, tp) {
		traverse(tile.S, tp, nil)
	}
	if canMove(tile.E, tp) {
		traverse(tile.E, tp, nil)
	}
	if canMove(tile.W, tp) {
		traverse(tile.W, tp, nil)
	}
	if canMove(tile.NE, tp) {
		traverse(tile.NE, tp, nil)
	}
	if canMove(tile.SE, tp) {
		traverse(tile.SE, tp, nil)
	}
	if canMove(tile.SW, tp) {
		traverse(tile.SW, tp, nil)
	}
	if canMove(tile.NW, tp) {
		traverse(tile.NW, tp, nil)
	}
}

//FindWords ...
func FindWords(board *board.Board, wc []wordcolumn.WordColumn) {
	wordColumn = wc
	var wg sync.WaitGroup

	wg.Add(board.Size)

	for i := 0; i < board.Size; i++ {
		go traverse(&board.Tiles[i], nil, &wg)
	}

	wg.Wait()
}
