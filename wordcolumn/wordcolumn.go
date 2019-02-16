package wordcolumn

//WordColumn ...
type WordColumn struct {
	RootIndex      uint16
	WordCount      uint16
	LongestWordLen uint16
	Words          map[string]struct{}
}

//AddWord ...
func (wc *WordColumn) AddWord(word string, rootIdx uint16) {
	if _, ok := wc.Words[word]; ok {
		return
	}
	wordLen := len(word)

	if wordLen < 3 {
		return
	}

	wc.RootIndex = uint16(rootIdx)
	wc.Words[word] = struct{}{}

	if wordLen > int(wc.LongestWordLen) {
		wc.LongestWordLen = uint16(wordLen)
	}

	wc.WordCount++
}
