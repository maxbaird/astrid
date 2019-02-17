package configuration

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
)

//Config ...
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

//Default values to be used if error occurs
const defaultMaxWordLength = 9
const defaultMinWordLength = 3
const defaultMaxWordsPerRow = 10
const defaultWordColumnsPerRow = 16
const defaultSortDescending = true
const defaultEnableHighlighting = true
const defaultHighlightLetters = "xqzjy"
const defaultLexisFilePath = "lexis"

//Other constants
const longestWordLength = 16
const maxWordColumnsPerRow = 16

//Config ...
var Config config

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
	if Config.MaxWordLength <= 0 || Config.MaxWordLength > longestWordLength {
		fmt.Fprintf(os.Stderr, "MaxWordLength must be between 0 and %d. Defaulting to %d.\n", longestWordLength, defaultMaxWordLength)
		Config.MaxWordLength = defaultMaxWordLength
	}

	if Config.MinWordLength <= 0 || Config.MinWordLength >= longestWordLength {
		fmt.Fprintf(os.Stderr, "MinWordLength must be between 0 and %d. Defaulting to %d.\n", longestWordLength, defaultMinWordLength)
		Config.MinWordLength = defaultMinWordLength
	}

	if Config.MinWordLength >= Config.MaxWordLength {
		fmt.Fprintf(os.Stderr, "MinWordLength must be less than MaxWordLength. Defaulting to %d.\n", defaultMinWordLength)
		Config.MinWordLength = defaultMinWordLength
	}

	if Config.MaxWordsPerRow <= 0 {
		fmt.Fprintf(os.Stderr, "MaxWordsPerRow cannot be less than 0. Defaulting to %d.\n", defaultMaxWordsPerRow)
		Config.MaxWordsPerRow = defaultMaxWordsPerRow
	}

	if Config.WordColumnsPerRow <= 0 || Config.WordColumnsPerRow > maxWordColumnsPerRow {
		fmt.Fprintf(os.Stderr, "WordColumnsPerRow must be between 0 and %d. Defaulting to %d.\n", maxWordColumnsPerRow,
			defaultWordColumnsPerRow)
		Config.WordColumnsPerRow = defaultWordColumnsPerRow
	}

	_, err := os.Stat(Config.LexisFilePath)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Lexis file \"%s\" not found. Defaulting to \"%s\".\n", Config.LexisFilePath,
			defaultLexisFilePath)

		Config.LexisFilePath = defaultLexisFilePath
	}
}

//ReadConfig ...
func ReadConfig() {
	_, err := os.Stat(configFile)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Config file missing: \"%s\". Using default configuration.", configFile)
	}

	if _, err := toml.DecodeFile(configFile, &Config); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading config file: %s.\nUsing default configuration.", err)
		Config = defaultConfig
		return
	}
	validateConfig()
}
