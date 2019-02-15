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
	letters      [][]rune
	traversePath []map[int]struct{}
}

func canMove(tile *tile.Tile, tilePath *path, depth int) bool {
	if tile == nil {
		return false
	}

	if _, ok := tilePath.traversePath[depth][int(tile.ID)]; ok {
		return false
	}

	return true
}

func traverse(tile *tile.Tile, p []path, pPathIdx int, depth int, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}

	if depth == 9 {
		return
	}

	idx := int(tile.ID) - 1

	if pPathIdx >= 0 {
		for k, v := range p[pPathIdx].traversePath[depth-1] {
			p[idx].traversePath[depth][k] = v
		}

		copy(p[idx].letters[idx], p[pPathIdx].letters[pPathIdx])
	}

	p[idx].traversePath[depth][int(tile.ID)] = struct{}{}
	//fmt.Printf("idx: %d, depth:%d len: %d\n", idx, depth, len(p[idx].letters))
	p[idx].letters[idx][depth] = tile.Letter
	//fmt.Println("After")

	word := string(p[idx].letters[idx][0 : depth+1])

	//fmt.Println(word)
	if lexis.IsWord(word) {
		fmt.Println(word)
	}

	if canMove(tile.N, &p[idx], depth) {
		traverse(tile.N, p, idx, depth+1, nil)
	}
	if canMove(tile.S, &p[idx], depth) {
		traverse(tile.S, p, idx, depth+1, nil)
	}
	if canMove(tile.E, &p[idx], depth) {
		traverse(tile.E, p, idx, depth+1, nil)
	}
	if canMove(tile.W, &p[idx], depth) {
		traverse(tile.W, p, idx, depth+1, nil)
	}
	if canMove(tile.NE, &p[idx], depth) {
		traverse(tile.NE, p, idx, depth+1, nil)
	}
	if canMove(tile.SE, &p[idx], depth) {
		traverse(tile.SE, p, idx, depth+1, nil)
	}
	if canMove(tile.SW, &p[idx], depth) {
		traverse(tile.SW, p, idx, depth+1, nil)
	}
	if canMove(tile.NW, &p[idx], depth) {
		traverse(tile.NW, p, idx, depth+1, nil)
	}
}

func doNothing(wg *sync.WaitGroup) {
	wg.Done()
}

//FindWords ...
func FindWords(board *board.Board) {
	var wg sync.WaitGroup

	p := make([]path, 16)

	fmt.Printf("Length of P: %d\n", len(p))

	for i := 0; i < 16; i++ {
		p[i].letters = make([][]rune, 16)
		p[i].traversePath = make([]map[int]struct{}, 16)

		for j := 0; j < 16; j++ {
			p[i].traversePath[j] = make(map[int]struct{})
			p[i].letters[j] = make([]rune, 9)
		}
	}

	//wg.Add(len(board.Tiles))
	wg.Add(1)
	p[0].root = int(board.Tiles[0].ID) - 1
	//fmt.Println("About to panic")

	//fmt.Println(p[0].letters)
	go traverse(&board.Tiles[0], p, -1, 0, &wg)
	//go doNothing(&wg)
	//for i, tile := range board.Tiles {
	//	p[i].root = int(tile.ID) - 1
	//	go traverse(&tile, &p[i], &wg)
	//}

	wg.Wait()
}
