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
		lexisTrie.Insert(reader.Text())
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
