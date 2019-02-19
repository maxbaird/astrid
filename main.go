//heh heh heh...
package main

import (
	"astrid/configuration"
	"astrid/lexis"
	"astrid/strid"
)

const height int = 4
const width int = 4

func main() {
	configuration.ReadConfig()
	lexis.LoadLexis()
	blitz := blitz.New(height, width)
	blitz.Start()
}
