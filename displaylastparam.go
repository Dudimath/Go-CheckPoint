// displaylastparam
// Instructions
// Write a program that displays its last argument, if there is one.

// Usage
// $ go run . hello there
// there
// $ go run . "hello there" how are you
// you
// $ go run . "hello there"
// hello there
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

	args := os.Args[1:]
	lastWord := args[len(args)-1]

	for _, c := range lastWord {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
