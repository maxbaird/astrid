package welcome

import (
	"fmt"
)

//PrintWelcome ...
func PrintWelcome() {
	welcome := `
   ▄████████    ▄████████     ███        ▄████████  ▄█  ████████▄
  ███    ███   ███    ███ ▀█████████▄   ███    ███ ███  ███   ▀███
  ███    ███   ███    █▀     ▀███▀▀██   ███    ███ ███▌ ███    ███
  ███    ███   ███            ███   ▀  ▄███▄▄▄▄██▀ ███▌ ███    ███
▀███████████ ▀███████████     ███     ▀▀███▀▀▀▀▀   ███▌ ███    ███
  ███    ███          ███     ███     ▀███████████ ███  ███    ███
  ███    ███    ▄█    ███     ███       ███    ███ ███  ███   ▄███
  ███    █▀   ▄████████▀     ▄████▀     ███    ███ █▀   ████████▀
                                        ███    ███
`
	fmt.Println(welcome)

	about := `
Astrid is a word finder for the currently popular Messenger game "Word Blitz"
on the Facebook platform. Enter the 16 letters of the word puzzle (row by row)
in one line and then hit enter. All possible words will then be displayed in
columns respective to the rank of each letter.
`
	fmt.Println(about)
}
