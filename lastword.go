// lastword
// Instructions
// Write a program that takes a string and displays its last word, followed by a newline ('\n').

// A word is a section of string delimited by spaces or by the start/end of the string.

// The output will be followed by a newline ('\n').

// If the number of arguments is different from 1, or if there are no word, the program displays nothing.

// Usage
// $ go run . "FOR PONY" | cat -e
// PONY$
// $ go run . "this        ...       is sparta, then again, maybe    not" | cat -e
// not$
// $ go run . "  "
// $ go run . "a" "b"
// $ go run . "  lorem,ipsum  " | cat -e
// lorem,ipsum$
// $

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) <= 1 || os.Args[1] == " " || len(os.Args) != 2 {
		return
	}
	args := os.Args[1]

	var indexes []int
	spaceIndex := 0
	lastSpaceIndex := 0

	for i, c := range args {
		if c == ' ' && i != len(args)-1 && i > spaceIndex {
			spaceIndex = i
		} else if c == ' ' && i == len(args)-1 && i > lastSpaceIndex {
			lastSpaceIndex = i
		}
	}

	for i, c := range args {
		if c == ' ' {
			indexes = append(indexes, i)
		}
	}

	var lastWord string
	for i, c := range args {
		if c == ' ' && i != len(args)-1 {
			lastWord = args[spaceIndex+1:]
		} else if c == ' ' && i == len(args)-1 {
			lastWord = args[len(indexes)-2 : lastSpaceIndex-1]
		}
	}

	for _, arg := range lastWord {
		z01.PrintRune(arg)
	}
	z01.PrintRune('\n')
}


// Stella solution

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 || os.Args[1] == "  " || os.Args[1] == " " {
		return
	}
	args := os.Args[1]
	lastWord := ""
	wordStart := false
	for i := len(args) - 1; i >= 0; i-- {
		if args[i] != ' ' {
			wordStart = true
			lastWord = string(args[i]) + lastWord
		} else if wordStart {
			break
		}
	}
	for _, r := range lastWord {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

// Stella Solution 2

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	args, lastword, start := os.Args[1], "", false

	for i := len(args) - 1; i >= 0; i-- {
		if args[i] != ' ' {
			start = true
			lastword = string(args[i]) + lastword
		} else if start {
			break
		}
	}
	if len(lastword) > 0 {
		for _, char := range lastword {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}
