// alphamirror
// Instructions
// Write a program called alphamirror that takes a string as argument and displays this string after replacing each alphabetical character with the opposite alphabetical character.

// The case of the letter remains unchanged, for example :

// 'a' becomes 'z', 'Z' becomes 'A' 'd' becomes 'w', 'M' becomes 'N'

// The final result will be followed by a newline ('\n').

// If the number of arguments is different from 1, the program prints a new line.

// Usage
// $ go run . "abc"
// zyx$
// $
// $ go run . "My horse is Amazing." | cat -e
// Nb slihv rh Znzarmt.$
// $
// $ go run .
// $

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) <= 1 {
		return
	}

	args := os.Args[1]
	var mirror string
	for _, c := range args {
		if c >= 'A' && c <= 'Z' {
			c = ('Z' - c) + 'A'
			mirror += string(c)
		} else if c >= 'a' && c <= 'z' {
			c = ('z' - c) + 'a'
			mirror += string(c)
		} else {
			mirror += string(c)
		}
	}
	for _, c := range mirror {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
