// displayfirstparam
// Instructions
// Write a program that displays its first argument, if there is one.

// Usage
// $ go run . hello there
// hello
// $ go run . "hello there" how are you
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
	firstWord := args[0]

	for _, arg := range firstWord {
		z01.PrintRune(arg)
	}
	z01.PrintRune('\n')
}
