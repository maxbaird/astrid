package printer

import (
	"astrid/board"
	"astrid/wordcolumn"
	"fmt"
	"sync"
)

type printColumn struct {
	words []string
}

var printColumns []printColumn

//PrintWords ...
func PrintWords(board *board.Board, wordColumn []wordcolumn.WordColumn) {
	printColumns = make([]printColumn, board.Size)

	var wg sync.WaitGroup
	wg.Add(int(board.Size))

	for i, wc := range wordColumn {
		go func(pc *printColumn, wc wordcolumn.WordColumn) {
			defer wg.Done()
			pc.words = make([]string, len(wc.Words))
			j := 0
			for k := range wc.Words {
				pc.words[j] = k
				j++
			}
		}(&printColumns[i], wc)
	}
	wg.Wait()

	for _, pc := range printColumns {
		fmt.Println(pc.words)
	}
}
