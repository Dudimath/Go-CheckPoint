// printreversealphabet
// Instructions
// Write a program that prints the Latin alphabet in lowercase in reverse order (from 'z' to 'a') on a single line.

// A line is a sequence of characters preceding the end of line character ('\n').

// Please note that casting is not allowed for this exercise!

// Usage
// $ go run .
// zyxwvutsrqponmlkjihgfedcba
// $

package main 

import (
	"github.com/01-edu/z01"
)

func main() {
	for i:= 'z'; i >= 'a'; i-- {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}