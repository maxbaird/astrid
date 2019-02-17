package configuration

import (
	"github.com/BurntSushi/toml"
	"log"
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

//Config ...
var Config config

//ReadConfig ...
func ReadConfig() {
	_, err := os.Stat(configFile)

	if err != nil {
		log.Fatal("Config file missing: ", configFile)
	}

	if _, err := toml.DecodeFile(configFile, &Config); err != nil {
		log.Fatal(err)
	}
}
