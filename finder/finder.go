package finder

import (
	"astrid/board"
	//"astrid/lexis"
	"astrid/tile"
	//"fmt"
	"sync"
)

var finderBoard *board.Board

type path struct {
	root         int
	traversePath []int
}

func canMove(tile *tile.Tile, tilePath *path) bool {
	if tile == nil {
		return false
	}

	for _, x := range tilePath.traversePath {
		if x+1 == int(tile.ID) {
			return false
		}
	}
	return true
}

func traverse(tile *tile.Tile, letters []rune, p *path, depth int, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}

	if depth == 9 {
		return
	}

	str := make([]rune, 9)
	tilePath := &path{}
	tilePath.traversePath = make([]int, finderBoard.GetBoardSize())

	//for i := range tilePath.traversePath {
	//	tilePath.traversePath[i] = -1
	//}

	if p == nil { //Will be nil for initial call
		tilePath.root = int(tile.ID)
	} else {
		copy(tilePath.traversePath, p.traversePath)
		tilePath.root = p.root
	}

	if letters != nil {
		copy(str, letters)
		str[depth] = tile.Letter
	} else {
		str[0] = tile.Letter //Handle initial traversal call
	}

	//word := string(str[0 : depth+1])

	//if lexis.IsWord(word) {
	//	fmt.Println(word)
	//}

	if canMove(tile.N, tilePath) {
		traverse(tile.N, str, tilePath, depth+1, nil)
	}
	if canMove(tile.S, tilePath) {
		traverse(tile.S, str, tilePath, depth+1, nil)
	}
	if canMove(tile.E, tilePath) {
		traverse(tile.E, str, tilePath, depth+1, nil)
	}
	if canMove(tile.W, tilePath) {
		traverse(tile.W, str, tilePath, depth+1, nil)
	}
	if canMove(tile.NE, tilePath) {
		traverse(tile.NE, str, tilePath, depth+1, nil)
	}
	if canMove(tile.SE, tilePath) {
		traverse(tile.SE, str, tilePath, depth+1, nil)
	}
	if canMove(tile.SW, tilePath) {
		traverse(tile.SW, str, tilePath, depth+1, nil)
	}
	if canMove(tile.NW, tilePath) {
		traverse(tile.NW, str, tilePath, depth+1, nil)
	}
}

//FindWords ...
func FindWords(board *board.Board) {
	finderBoard = board

	var i uint16
	var wg sync.WaitGroup

	for i = 0; i < finderBoard.GetBoardSize(); i++ {
		wg.Add(1)
		go traverse(&finderBoard.Tiles[i], nil, nil, 0, &wg)
	}
	wg.Wait()
}
