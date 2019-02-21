package wordcolumn

import (
	"github.com/maxbaird/astrid/config"
)

//WordColumn ...
type WordColumn struct {
	RootIndex      int
	WordCount      int
	LongestWordLen int
	Words          map[string]struct{}
}

//AddWord ...
func (wc *WordColumn) AddWord(word string, rootIdx int) {
	if _, ok := wc.Words[word]; ok {
		return
	}
	wordLen := len(word)

	if wordLen < config.MinWordLength {
		return
	}

	wc.RootIndex = rootIdx
	wc.Words[word] = struct{}{}

	if wordLen > wc.LongestWordLen {
		wc.LongestWordLen = wordLen
	}

	wc.WordCount++
}
