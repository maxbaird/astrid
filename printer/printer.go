package printer

import (
	"astrid/board"
	"astrid/wordcolumn"
	"fmt"
	"sort"
	"sync"
)

type byLength []string

func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return len(s[i]) > len(s[j])
}

type printColumn struct {
	idx            uint16
	words          []string
	longestWordLen uint16
}

const spaceBetweenColumns uint16 = 2

var printColumns []printColumn
var longestWordLen uint16

func makePrintColumns(board *board.Board, wordColumns []wordcolumn.WordColumn) {
	printColumns = make([]printColumn, board.Size)

	var wg sync.WaitGroup

	wg.Add(int(board.Size))

	for i, wc := range wordColumns {
		go func(pc *printColumn, wc wordcolumn.WordColumn) {
			defer wg.Done()

			pc.words = make([]string, len(wc.Words))
			pc.longestWordLen = wc.LongestWordLen
			pc.idx = wc.RootIndex
			j := 0

			for k := range wc.Words {
				pc.words[j] = k
				j++
			}
			sort.Sort(byLength(pc.words))
		}(&printColumns[i], wc)
	}
	wg.Wait()
}

func findLongestWord() {
	for _, pc := range printColumns {
		if pc.longestWordLen > longestWordLen {
			longestWordLen = pc.longestWordLen
		}
	}
}

func pad(length uint16) uint16 {
	return (longestWordLen - length) + spaceBetweenColumns
}

func printColumnHeaders(start int, end int) {
	for i := start; i < end; i++ {
		str := fmt.Sprintf("[%d]", i+1)
		//str := fmt.Sprintf("%s%c%d%c", s1, '%', pad(uint16(len(s1))), 's')
		fmt.Printf("%s%*s", str, pad(uint16(len(str))), "")
	}
	fmt.Println()
}

//PrintWords ...
func PrintWords(board *board.Board, wordColumns []wordcolumn.WordColumn) {
	makePrintColumns(board, wordColumns)
	findLongestWord()

	colHeaderStart := 0
	colHeaderEnd := colHeaderStart + 16

	var i uint16
	for i = 0; i < board.Size; i = i + 16 {
		printColumnHeaders(colHeaderStart, colHeaderEnd)

		colHeaderStart += 16
		colHeaderEnd = colHeaderStart + 16
		if colHeaderEnd >= int(board.Size) {
			colHeaderEnd = int(board.Size)
		}
	}
}
