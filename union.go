// union
// Instructions
// Write a program that takes two string and displays, without doubles, the characters that appear in either one of the string.

// The display will be in the same order that the characters appear on the command line and will be followed by a newline ('\n').

// If the number of arguments is different from 2, then the program displays a newline ('\n').

// Usage
// $ go run . zpadinton paqefwtdjetyiytjneytjoeyjnejeyj | cat -e
// zpadintoqefwjy$
// $
// $ go run . ddf6vewg64f gtwthgdwthdwfteewhrtag6h4ffdhsd | cat -e
// df6vewg4thras$
// $
// $ go run . rien "cette phrase ne cache rien" | cat -e
// rienct phas$
// $
// $ go run . | cat -e
// $
// $ go run . rien | cat -e
// $

package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	args := os.Args[1:]

	union := make(map[rune]bool)

	for _, c := range args[0] + args[1] {
		if !union[c] { // checks if the letter is not in the union map yet 
			z01.PrintRune(c)
			union[c] = true // ones the letter is added the value becomes true
		}
	}
	z01.PrintRune('\n')
}
