package strid

import (
	"github.com/maxbaird/astrid/board"
	"github.com/maxbaird/astrid/finder"
	"github.com/maxbaird/astrid/printer"
	"github.com/maxbaird/astrid/tile"
	"github.com/maxbaird/astrid/welcome"
	"github.com/maxbaird/astrid/wordcolumn"
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

//Strid ...
type Strid struct {
	WordColumn []wordcolumn.WordColumn
	Board      *board.Board
}

const prompt = ">> "

var maxLetters int

//New ...
func New(height int, width int) *Strid {
	tiles := make([]tile.Tile, height*width)
	board := board.New(tiles, height, width)

	maxLetters = height * width

	wc := make([]wordcolumn.WordColumn, board.Size)

	for i := 0; i < board.Size; i++ {
		wc[i].Words = make(map[string]struct{})
	}

	return &Strid{wc, board}
}

func validateInput(letters string) (string, bool) {
	letters = strings.Replace(letters, "\n", "", -1)
	letters = strings.Replace(letters, " ", "", -1)
	letters = strings.ToLower(letters)

	letterLen := len(letters)

	if letterLen == 0 {
		return letters, false
	}

	if letters == "test" {
		letters = "abcdefghijklmnop"
		return letters, true
	}

	if letterLen < maxLetters {
		fmt.Fprintf(os.Stderr, "%d letters needed; %d entered!\n", maxLetters, letterLen)
		return letters, false
	}

	f := func(r rune) bool {
		return r < 'a' || r > 'z'
	}

	if strings.IndexFunc(letters, f) != -1 {
		fmt.Fprintf(os.Stderr, "Only letters between a - z allowed.\n")
		return letters, false
	}

	return letters, true
}

func reset(strid *Strid) {
	strid.WordColumn = nil
	strid.WordColumn = make([]wordcolumn.WordColumn, strid.Board.Size)

	for i := 0; i < strid.Board.Size; i++ {
		strid.WordColumn[i].Words = make(map[string]struct{})
	}
}

func handleExit() {
	gracefulStop := make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	signal.Notify(gracefulStop, syscall.SIGHUP)
	signal.Notify(gracefulStop, syscall.SIGQUIT)

	go func() {
		<-gracefulStop
		fmt.Println("\nBye!")
		os.Exit(0)
	}()
}

//Start ...
func (strid *Strid) Start() {
	handleExit()

	welcome.PrintWelcome()
	fmt.Println("ctrl + C to exit.\n")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')

		if letters, ok := validateInput(input); ok {
			strid.Board.PlaceLetters(letters)
			finder.FindWords(strid.Board, strid.WordColumn)
			printer.PrintWords(strid.Board, strid.WordColumn)
			reset(strid)
		}
	}
}
