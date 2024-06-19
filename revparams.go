// revparams
// Instructions
// Write a program that prints the arguments received in the command line in reverse order.

// Example of output :

// $ go run . choumi is the best cat
// cat
// best
// the
// is
// choumi
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

	for i := len(args) - 1; i >= 0; i-- {
		param := args[i]
		for _, c := range param {
			z01.PrintRune(c)
		}
		z01.PrintRune('\n')
	}
}
