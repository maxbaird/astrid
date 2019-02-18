package blitz

import (
	"astrid/board"
	"astrid/finder"
	"astrid/printer"
	"astrid/tile"
	"astrid/welcome"
	"astrid/wordcolumn"
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

//Blitz ...
type Blitz struct {
	WordColumn []wordcolumn.WordColumn
	Board      *board.Board
}

const prompt = ">> "

var maxLetters int

//New ...
func New(height int, width int) *Blitz {
	tiles := make([]tile.Tile, height*width)
	board := board.New(tiles, height, width)

	maxLetters = height * width

	wc := make([]wordcolumn.WordColumn, board.Size)

	for i := 0; i < board.Size; i++ {
		wc[i].Words = make(map[string]struct{})
	}

	return &Blitz{wc, board}
}

func validateInput(letters string) (string, bool) {
	letters = strings.Replace(letters, "\n", "", -1)
	letters = strings.Replace(letters, " ", "", -1)
	letters = strings.ToLower(letters)

	if letters == "test" {
		letters = "abcdefghijklmnop"
		return letters, true
	}

	letterLen := len(letters)

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

func reset(blitz *Blitz) {
	blitz.WordColumn = nil
	blitz.WordColumn = make([]wordcolumn.WordColumn, blitz.Board.Size)

	for i := 0; i < blitz.Board.Size; i++ {
		blitz.WordColumn[i].Words = make(map[string]struct{})
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
func (blitz *Blitz) Start() {
	handleExit()

	welcome.PrintWelcome()
	fmt.Println("ctrl + C to exit.\n")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')

		if letters, ok := validateInput(input); ok {
			blitz.Board.PlaceLetters(letters)
			finder.FindWords(blitz.Board, blitz.WordColumn)
			printer.PrintWords(blitz.Board, blitz.WordColumn)
			reset(blitz)
		}
	}
}
