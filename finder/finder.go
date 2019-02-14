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
	traversePath []map[int]struct{}
}

func canMove(tile *tile.Tile, tilePath *path) bool {
	if tile == nil {
		return false
	}

	//if tilePath.depth == 9 { //This may not be necessary
	//	return false
	//}

	//fmt.Printf("Tile id: %d\n", int(tile.ID))
	if _, ok := tilePath.traversePath[tilePath.depth][int(tile.ID)]; ok {
		return false
	}

	return true
}

func traverse(tile *tile.Tile, p *path, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}

	if p.depth == 9 {
		return
	}

	//str := make([]rune, 9)
	//tilePath := &path{}
	//tilePath.traversePath = make(map[int]struct{})

	//if p == nil { //Will be nil for initial call
	//	tilePath.root = int(tile.ID)
	//} else {
	//for k, v := range p.traversePath { //Copy the map
	//	tilePath.traversePath[k] = v
	//}
	//tilePath.root = p.root
	//}

	//if letters != nil {
	//copy(str, letters)
	//fmt.Println(p.root)
	p.traversePath[p.depth][int(tile.ID)] = struct{}{}
	//fmt.Println(p.traversePath)
	p.letters[p.depth] = tile.Letter
	p.depth++
	//}

	word := string(p.letters[0:p.depth])

	//fmt.Println(word)
	if lexis.IsWord(word) {
		fmt.Println(word)
	}

	if canMove(tile.N, p) {
		traverse(tile.N, p, nil)
	}
	if canMove(tile.S, p) {
		traverse(tile.S, p, nil)
	}
	if canMove(tile.E, p) {
		traverse(tile.E, p, nil)
	}
	if canMove(tile.W, p) {
		traverse(tile.W, p, nil)
	}
	if canMove(tile.NE, p) {
		traverse(tile.NE, p, nil)
	}
	if canMove(tile.SE, p) {
		traverse(tile.SE, p, nil)
	}
	if canMove(tile.SW, p) {
		traverse(tile.SW, p, nil)
	}
	if canMove(tile.NW, p) {
		traverse(tile.NW, p, nil)
	}
}

//FindWords ...
func FindWords(board *board.Board) {
	var wg sync.WaitGroup

	p := make([]path, 16)

	fmt.Printf("Length of P: %d\n", len(p))

	for i := 0; i < 16; i++ {
		p[i].letters = make([]rune, 9)
		p[i].traversePath = make([]map[int]struct{}, 9)

		for j := 0; j < 9; j++ {
			p[i].traversePath[j] = make(map[int]struct{})
		}
	}

	wg.Add(len(board.Tiles))
	//wg.Add(1)
	//p[0].root = int(board.Tiles[0].ID) - 1
	//go traverse(&board.Tiles[0], &p[0], &wg)

	for i, tile := range board.Tiles {
		p[i].root = int(tile.ID) - 1
		go traverse(&tile, &p[i], &wg)
	}

	wg.Wait()
}
