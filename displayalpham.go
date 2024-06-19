// displayalpham
// Instructions
// Write a program that displays the alphabet, with even letters in uppercase, and odd letters in lowercase, followed by a newline ('\n').

// Usage
// $ go run . | cat -e
// aBcDeFgHiJkLmNoPqRsTuVwXyZ$
// $

package main

import "github.com/01-edu/z01"

func main() {
	for i := 'a'; i <= 'z'; i++ {
		if i%2 == 0 {
			c := i - 32
			z01.PrintRune(c)
		} else {
			z01.PrintRune(i)
		}
	}
	z01.PrintRune('\n')
}
