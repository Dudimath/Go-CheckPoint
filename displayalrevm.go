// displayalrevm
// Instructions
// Write a program that displays the alphabet in reverse, with even letters in uppercase, and odd letters in lowercase, followed by a newline ('\n').

// Usage
// $ go run . | cat -e
// zYxWvUtSrQpOnMlKjIhGfEdCbA$
// $

package main

import "github.com/01-edu/z01"

func main() {
	for i := 'Z'; i >= 'A'; i-- {
		if i%2 == 0 {
			c := i + 32
			z01.PrintRune(c)
		} else {
			z01.PrintRune(i)
		}
	}
	z01.PrintRune('\n')
}
