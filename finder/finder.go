package finder

import (
	"astrid/board"
	"astrid/lexis"
	"astrid/tile"
	"fmt"
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

	var i uint16

	for i = 0; i < finderBoard.GetBoardSize(); i++ {
		if tilePath.traversePath[i] == int(tile.ID) {
			return false
		}
	}
	return true
}

func traverse(tile *tile.Tile, letters []rune, p *path, depth int) {
	if depth == 16 {
		return
	}

	str := make([]rune, 16)
	tilePath := &path{}
	tilePath.traversePath = make([]int, finderBoard.GetBoardSize())

	for i := range tilePath.traversePath {
		tilePath.traversePath[i] = -1
	}

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

	word := string(str[0 : depth+1])

	if lexis.IsWord(word) {
		fmt.Println(word)
	}

	if canMove(tile.N, tilePath) {
		traverse(tile.N, str, tilePath, depth+1)
	}
	if canMove(tile.S, tilePath) {
		traverse(tile.S, str, tilePath, depth+1)
	}
	if canMove(tile.E, tilePath) {
		traverse(tile.E, str, tilePath, depth+1)
	}
	if canMove(tile.W, tilePath) {
		traverse(tile.W, str, tilePath, depth+1)
	}
	if canMove(tile.NE, tilePath) {
		traverse(tile.NE, str, tilePath, depth+1)
	}
	if canMove(tile.SE, tilePath) {
		traverse(tile.SE, str, tilePath, depth+1)
	}
	if canMove(tile.SW, tilePath) {
		traverse(tile.SW, str, tilePath, depth+1)
	}
	if canMove(tile.NW, tilePath) {
		traverse(tile.NW, str, tilePath, depth+1)
	}
}

//FindWords ...
func FindWords(board *board.Board) {
	finderBoard = board

	var i uint16

	for i = 0; i < finderBoard.GetBoardSize(); i++ {
		traverse(&finderBoard.Tiles[i], nil, nil, 0)
	}
}
