// hello
// Instructions
// Write a program that displays "Hello World!" followed by a newline ('\n').

// Usage
// $ go run .
// Hello World!
// $

package main 

import "github.com/01-edu/z01"

func main() {
	s := "Hello World!"
	for _, c := range s {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}