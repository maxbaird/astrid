//heh heh heh...
package main

import (
	"github.com/maxbaird/astrid/config"
	"github.com/maxbaird/astrid/lexis"
	"github.com/maxbaird/astrid/strid"
)

const height int = 4
const width int = 4

func main() {
	config.ReadConfig()
	lexis.LoadLexis()
	strid := strid.New(height, width)
	strid.Start()
}
