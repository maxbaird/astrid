package trie

const alphabetSize uint16 = 26

//Trie ...
type Trie struct {
	children    [alphabetSize]*Trie
	isEndOfWord bool
}

func charToIndex(c rune) uint {
	return uint(c) - uint('a')
}

//Insert ...
func (trie *Trie) Insert(letters string) {
	length := len(letters)

	for level := 0; level < length; level++ {
		index := charToIndex(rune(letters[level]))

		if trie.children[index] == nil {
			trie.children[index] = &Trie{}
		}
		trie = trie.children[index]
	}
	trie.isEndOfWord = true
}

//Has ...
func (trie *Trie) Has(key string) bool {
	length := len(key)

	for level := 0; level < length; level++ {
		index := charToIndex(rune(key[level]))

		if trie.children[index] == nil {
			return false
		}
		trie = trie.children[index]
	}

	return (trie != nil && trie.isEndOfWord)
}
