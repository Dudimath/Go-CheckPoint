package main

import "fmt"

func SplitWhiteSpaces(s string) []string {
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

func main() {
	fmt.Printf("%#v\n", SplitWhiteSpaces("Hello how are you?"))
}
