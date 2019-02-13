package lexis

import (
	"astrid/trie"
	"bufio"
	"log"
	"os"
)

var lexisTrie trie.Trie

//LoadLexis ...
func LoadLexis() {
	file, err := os.Open("wordList")

	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewScanner(file)

	for reader.Scan() {
		txt := reader.Text()
		if len(txt) <= 9 {
			lexisTrie.Insert(txt)
		}
	}

	if reader.Err() != nil {
		log.Fatal(err)
	}

	file.Close()
}

//IsWord ...
func IsWord(letters string) bool {
	return lexisTrie.Has(letters)
}
