package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
)

//Default values to be used if error occurs
const defaultMaxWordLength = 9
const defaultMinWordLength = 3
const defaultMaxWordsPerRow = 10
const defaultWordColumnsPerRow = 16
const defaultSortDescending = true
const defaultEnableHighlighting = true
const defaultHighlightLetters = "xqzjy"
const defaultLexisFilePath = "wordList"

var (
	//MaxWordLength ...
	MaxWordLength = defaultMaxWordLength
	//MinWordLength ...
	MinWordLength = defaultMinWordLength
	//MaxWordsPerRow ..
	MaxWordsPerRow = defaultMaxWordsPerRow
	//WordColumnsPerRow ..
	WordColumnsPerRow = defaultWordColumnsPerRow
	//SortDescending ...
	SortDescending = defaultSortDescending
	//EnableHighlighting ..
	EnableHighlighting = defaultEnableHighlighting
	//HighlightLetters ...
	HighlightLetters = defaultHighlightLetters
	//LexisFilePath ...
	LexisFilePath = defaultLexisFilePath
)

type config struct {
	MaxWordLength      int
	MinWordLength      int
	MaxWordsPerRow     int
	WordColumnsPerRow  int
	SortDescending     bool
	EnableHighlighting bool
	HighlightLetters   string
	LexisFilePath      string
}

const configFile = "config.conf"

//Other constants
const longestWordLength = 16
const maxWordColumnsPerRow = 16

var loadedConfig config

var defaultConfig = config{defaultMaxWordLength,
	defaultMinWordLength,
	defaultMaxWordsPerRow,
	defaultWordColumnsPerRow,
	defaultSortDescending,
	defaultEnableHighlighting,
	defaultHighlightLetters,
	defaultLexisFilePath,
}

func validateConfig() {
	if loadedConfig.MaxWordLength <= 0 || loadedConfig.MaxWordLength > longestWordLength {
		fmt.Fprintf(os.Stderr, "MaxWordLength must be between 0 and %d. Defaulting to %d.\n", longestWordLength, defaultMaxWordLength)
		MaxWordLength = defaultMaxWordLength
	} else {
		MaxWordLength = loadedConfig.MaxWordLength
	}

	if loadedConfig.MinWordLength <= 0 || loadedConfig.MinWordLength >= longestWordLength {
		fmt.Fprintf(os.Stderr, "MinWordLength must be between 0 and %d. Defaulting to %d.\n", longestWordLength, defaultMinWordLength)
		MinWordLength = defaultMinWordLength
	} else {
		MinWordLength = loadedConfig.MinWordLength
	}

	if loadedConfig.MinWordLength >= loadedConfig.MaxWordLength {
		fmt.Fprintf(os.Stderr, "MinWordLength must be less than MaxWordLength. Defaulting to %d.\n", defaultMinWordLength)
		MinWordLength = defaultMinWordLength
	} else {
		MinWordLength = loadedConfig.MinWordLength
	}

	if loadedConfig.MaxWordsPerRow <= 0 {
		fmt.Fprintf(os.Stderr, "MaxWordsPerRow cannot be less than 0. Defaulting to %d.\n", defaultMaxWordsPerRow)
		MaxWordsPerRow = defaultMaxWordsPerRow
	} else {
		MaxWordsPerRow = loadedConfig.MaxWordsPerRow
	}

	if loadedConfig.WordColumnsPerRow <= 0 || loadedConfig.WordColumnsPerRow > maxWordColumnsPerRow {
		fmt.Fprintf(os.Stderr, "WordColumnsPerRow must be between 0 and %d. Defaulting to %d.\n", maxWordColumnsPerRow,
			defaultWordColumnsPerRow)
		WordColumnsPerRow = defaultWordColumnsPerRow
	} else {
		WordColumnsPerRow = loadedConfig.WordColumnsPerRow
	}

	file, err := os.Stat(loadedConfig.LexisFilePath)

	if err != nil || !file.Mode().IsRegular() {
		fmt.Fprintf(os.Stderr, "Lexis file \"%s\" not found. Defaulting to \"%s\".\n", loadedConfig.LexisFilePath,
			defaultLexisFilePath)

		LexisFilePath = defaultLexisFilePath
	} else {
		LexisFilePath = loadedConfig.LexisFilePath
	}
}

//ReadConfig ...
func ReadConfig() {
	_, err := os.Stat(configFile)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Config file missing: \"%s\". Using default configuration.", configFile)
	}

	if _, err := toml.DecodeFile(configFile, &loadedConfig); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading config file: %s.\nUsing default configuration.", err)
		loadedConfig = defaultConfig
		return
	}
	validateConfig()
}
