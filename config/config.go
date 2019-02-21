package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
)

const (
	//Default values to be used if error occurs
	defaultMaxWordLength      = 9
	defaultMinWordLength      = 3
	defaultMaxWordsPerRow     = 10
	defaultWordColumnsPerRow  = 16
	defaultSortDescending     = true
	defaultEnableHighlighting = true
	defaultHighlightLetters   = "xqzjy"
	defaultLexisFilePath      = "wordList"

	//Constants not related to config values
	configFile           = "config.conf"
	longestWordLength    = 16
	maxWordColumnsPerRow = 16
)

//Configuration values needed by other packages
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
	} else {
		MaxWordLength = loadedConfig.MaxWordLength
	}

	if loadedConfig.MinWordLength <= 0 || loadedConfig.MinWordLength >= longestWordLength {
		fmt.Fprintf(os.Stderr, "MinWordLength must be between 0 and %d. Defaulting to %d.\n", longestWordLength, defaultMinWordLength)
	} else {
		MinWordLength = loadedConfig.MinWordLength
	}

	if loadedConfig.MinWordLength >= loadedConfig.MaxWordLength {
		fmt.Fprintf(os.Stderr, "MinWordLength must be less than MaxWordLength. Defaulting to %d.\n", defaultMinWordLength)
	} else {
		MinWordLength = loadedConfig.MinWordLength
	}

	if loadedConfig.MaxWordsPerRow <= 0 {
		fmt.Fprintf(os.Stderr, "MaxWordsPerRow cannot be less than 0. Defaulting to %d.\n", defaultMaxWordsPerRow)
	} else {
		MaxWordsPerRow = loadedConfig.MaxWordsPerRow
	}

	if loadedConfig.WordColumnsPerRow <= 0 || loadedConfig.WordColumnsPerRow > maxWordColumnsPerRow {
		fmt.Fprintf(os.Stderr, "WordColumnsPerRow must be between 0 and %d. Defaulting to %d.\n", maxWordColumnsPerRow,
			defaultWordColumnsPerRow)
	} else {
		WordColumnsPerRow = loadedConfig.WordColumnsPerRow
	}

	file, err := os.Stat(loadedConfig.LexisFilePath)

	if err != nil || !file.Mode().IsRegular() {
		fmt.Fprintf(os.Stderr, "Lexis file \"%s\" not found. Defaulting to \"%s\".\n", loadedConfig.LexisFilePath,
			defaultLexisFilePath)
	} else {
		LexisFilePath = loadedConfig.LexisFilePath
	}
}

//ReadConfig ...
func ReadConfig() {
	file, err := os.Stat(configFile)

	if err != nil || !file.Mode().IsRegular() {
		fmt.Fprintf(os.Stderr, "Config file missing: \"%s\". Using default configuration.", configFile)
		loadedConfig = defaultConfig
	} else if _, err := toml.DecodeFile(configFile, &loadedConfig); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading config file: %s.\nUsing default configuration.", err)
		loadedConfig = defaultConfig
	}

	validateConfig()
}
