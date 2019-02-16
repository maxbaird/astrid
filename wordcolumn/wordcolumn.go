package wordcolumn

//WordColumn ...
type WordColumn struct {
	rootIndex      uint16
	wordCount      uint16
	longestWordLen uint16
	Words          map[string]struct{}
}

//AddWord ...
func (wc *WordColumn) AddWord(word string) {
	if _, ok := wc.Words[word]; ok {
		return
	}
	wordLen := len(word)

	if wordLen < 3 {
		return
	}

	wc.Words[word] = struct{}{}

	if wordLen > int(wc.longestWordLen) {
		wc.longestWordLen = uint16(wordLen)
	}

	wc.wordCount++
}
