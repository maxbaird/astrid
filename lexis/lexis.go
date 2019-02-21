package lexis

import (
	"bufio"
	"github.com/maxbaird/astrid/config"
	"github.com/maxbaird/astrid/trie"
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

	file, err := os.Open(config.LexisFilePath)

	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewScanner(file)

	for reader.Scan() {
		txt := reader.Text()
		if len(txt) <= config.MaxWordLength {
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
