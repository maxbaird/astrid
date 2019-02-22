package printer

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/maxbaird/astrid/board"
	"github.com/maxbaird/astrid/config"
	"github.com/maxbaird/astrid/wordcolumn"
	"sort"
	"strings"
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
	idx            int
	words          []string
	wordCount      int
	longestWordLen int
}

const spaceBetweenColumns int = 2

var printColumns []printColumn
var longestWordLen int

func makePrintColumns(board *board.Board, wordColumns []wordcolumn.WordColumn) {
	printColumns = make([]printColumn, board.Size)

	var wg sync.WaitGroup

	wg.Add(board.Size)

	for i, wc := range wordColumns {
		go func(pc *printColumn, wc wordcolumn.WordColumn) {
			defer wg.Done()

			pc.words = make([]string, len(wc.Words))
			pc.longestWordLen = wc.LongestWordLen
			pc.idx = wc.RootIndex
			pc.wordCount = wc.WordCount
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

func pad(pc printColumn, length int) int {
	return (pc.longestWordLen - length) + spaceBetweenColumns
}

func printColumnHeaders(start int, end int) {
	c := color.New(color.FgCyan, color.Bold)

	for i := start; i < end; i++ {
		str := fmt.Sprintf("[%d]", i+1)
		c.Printf("%s%*s", str, pad(printColumns[i], len(str)), "")
	}
	fmt.Println()
}

func getLongestColumn(start int, end int) int {
	var count int

	for i := start; i < end; i++ {
		if printColumns[i].wordCount > count {
			count = printColumns[i].wordCount
		}
	}

	return count
}

func printWord(pc printColumn, wordIdx int, endColumn bool) {
	var padding int
	word := pc.words[wordIdx]

	if !endColumn {
		padding = pad(pc, len(word))
	}

	str := fmt.Sprintf("%s%*s", word, padding, "")

	if strings.ContainsAny(word, config.HighlightLetters) {
		c := color.New(color.FgRed, color.Bold)
		c.Printf(str)
	} else {
		fmt.Printf(str)
	}
}

//PrintWords ...
func PrintWords(board *board.Board, wordColumns []wordcolumn.WordColumn) {
	makePrintColumns(board, wordColumns)
	findLongestWord()

	colsPerRow := config.WordColumnsPerRow
	maxWordsPerRow := config.MaxWordsPerRow

	colHeaderStart := 0
	colHeaderEnd := colHeaderStart + colsPerRow

	for i := 0; i < board.Size; i += colsPerRow {
		printColumnHeaders(colHeaderStart, colHeaderEnd)

		longestColumn := getLongestColumn(colHeaderStart, colHeaderEnd)
		numPrintedRows := 0

		for j := 0; j < longestColumn; j++ {
			if numPrintedRows == maxWordsPerRow {
				numPrintedRows = 0
				fmt.Println()
				printColumnHeaders(colHeaderStart, colHeaderEnd)
			}
			numPrintedRows++

			for k := colHeaderStart; k < colHeaderEnd; k++ {
				if printColumns[k].wordCount > j {
					printWord(printColumns[k], j, k == (colHeaderEnd-1))
				} else { //If there are no words to print
					fmt.Printf("%*s", pad(printColumns[k], 0), "")
				}
			}
			fmt.Println()
		}
		fmt.Println()

		colHeaderStart += colsPerRow
		colHeaderEnd = colHeaderStart + colsPerRow
		if colHeaderEnd >= board.Size {
			colHeaderEnd = board.Size
		}
	}

	for i := 0; i < (longestWordLen+spaceBetweenColumns)*colsPerRow; i++ {
		fmt.Print("+")
	}
	fmt.Printf("\n\n")
}
