package lexis

import (
	"astrid/trie"
	"bufio"
	"fmt"
	"log"
	"os"
)

//LoadLexis ...
func LoadLexis() {
	file, err := os.Open("wordList")

	if err != nil {
		log.Fatal(err)
	}

	trie := trie.New()

	fs := bufio.NewScanner(file)

	for fs.Scan() {
		txt := fs.Text()
		trie.Insert(txt)
	}

	s := "asshat"
	if trie.Has(s) {
		fmt.Printf("Found %s!\n", s)
	} else {
		fmt.Printf("%s not found :-(\n", s)
	}
	file.Close()
}

//IsWord ...
func IsWord(letters string) bool {
	fmt.Print(letters)
	return true
}
