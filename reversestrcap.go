// reversestrcap
// Instructions
// Write a program that takes one or more arguments and that, for each argument, puts the last letter of each word in uppercase and the rest in lowercase. It displays the result followed by a newline ('\n').

// If there are no argument, the program displays nothing.

// Usage
// $ go run . "First SMALL TesT" | cat -e
// firsT smalL tesT$
// $ go run . "SEconD Test IS a LItTLE EasIEr" "bEwaRe IT'S NoT HARd WhEN " " Go a dernier 0123456789 for the road e" | cat -e
// seconD tesT iS A littlE easieR$
// bewarE it'S noT harD wheN $
//  gO A dernieR 0123456789 foR thE roaD E$
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
	var words string
	args := os.Args[1]
	nwSlice := splitWhiteSpaces(args)

	for _, slice := range nwSlice {
		for i, c := range slice {
			if i != len(slice)-1 && c >= 'A' && c <= 'Z' {
				c = c + 32
				words += string(c)
			} else if i == len(slice)-1 && c >= 'a' && c <= 'z' {
				c = c - 32
				words += string(c)
			} else {
				words += string(c)
			}
		}
	}
	nStr := ""
	for _, c := range words {
		if len(nStr) < len(words) && c >= 'A' && c <= 'Z' {
			nStr += string(c) + " "
		} else {
			nStr += string(c)
		}
	}
	for _, c := range nStr {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}

func splitWhiteSpaces(s string) []string {
	wordSlice := []string{}
	var indx int
	for i, c := range s {
		if c == ' ' {
			if indx < i {
				wordSlice = append(wordSlice, s[indx:i])
			}
			indx = i + 1
		}
	}
	if indx < len(s) {
		wordSlice = append(wordSlice, s[indx:])
	}

	return wordSlice
}
