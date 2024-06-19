// revwstr
// Instructions
// Write a program that takes a string as a parameter, and prints its words in reverse, followed by a newline.

// A word is a sequence of alphanumerical characters.

// If the number of arguments is different from 1, the program will display nothing.

// In the parameters that are going to be tested, there will not be any extra spaces. (meaning that there will not be additional spaces at the beginning or at the end of the string and that words will always be separated by exactly one space).

// Usage
// $ go run . "the time of contempt precedes that of indifference"
// indifference of that precedes contempt of time the
// $ go run . "abcdefghijklm"
// abcdefghijklm
// $ go run . "he stared at the mountain"
// mountain the at stared he
// $ go run . "" | cat-e
// $
// $

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	if len(os.Args) != 2 || os.Args[1] == "" {
		return
	}

	args := os.Args[1]

	var start int
	var current int

	var words []string
	for i, c := range args {
		if c == ' ' {
			current = i
			words = append(words, args[start:current])
			start = current + 1
		}
	}

	words = append(words, args[start:])
	var	revSlice []string

	for i := len(words) - 1; i >= 0; i-- {
		revSlice = append(revSlice, words[i])
	}

	var revStr string

	for i, c := range revSlice {
		if i < len(revSlice) {
			revStr += c + " "
		} else {
			revStr += c
		}
	}
	for _, s := range revStr {
		z01.PrintRune(s)
	}
	z01.PrintRune('\n')
}
