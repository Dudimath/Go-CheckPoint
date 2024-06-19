// rot13
// Instructions
// Write a program that takes a string and displays it, replacing each of its letters by the letter 13 spots ahead in alphabetical order.

// 'z' becomes 'm' and 'Z' becomes 'M'. The case of the letter stays the same.

// The output will be followed by a newline ('\n').

// If the number of arguments is different from 1, the program displays nothing.

// Usage
// $ go run . "abc"
// nop
// $ go run . "hello there"
// uryyb gurer
// $ go run . "HELLO, HELP"
// URYYB, URYC
// $ go run .
// $

package main 

import (
	"os"
	"github.com/01-edu/z01"
)

const rot13 = 13

func main() {

	if len(os.Args) < 2 {
		return
	}
	args := os.Args[1]
	var words string

	for _, c := range args {
		if c >= 'a' && c <= 'z' && c + rot13 > 'z' {
			c = (c + rot13) - 26
			words += string(c)
		} else if c >= 'A' && c <= 'Z' && c + rot13 > 'Z' {
			c = (c + rot13) - 26
			words += string(c)
		} else if c >= 'a' && c <= 'z' && c + rot13 < 'z' {
			c = c + rot13
			words += string(c)
		} else if c >= 'A' && c <= 'Z' && c + rot13 < 'Z' {
			c = c + rot13
			words += string(c)
		} else {
			words += string(c)
		}
	}

	for _, c := range words {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}