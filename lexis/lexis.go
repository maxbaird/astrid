package lexis

import (
	"astrid/trie"
	"bufio"
	"log"
	"os"
)

var lexisTrie trie.Trie
var initialized bool

//LoadLexis ...
func LoadLexis() {
	if initialized {
		return
	}

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
	initialized = true
}

//IsWord ...
func IsWord(letters string) bool {
	if !initialized {
		log.Fatal("Lexis not loaded!")
	}
	return lexisTrie.Has(letters)
}
