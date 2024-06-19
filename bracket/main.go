package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		return
	}
	arg := os.Args[1:]

	for _, word := range arg {
		if IsCorrectBracket(word) || word == "" {
			fmt.Println("OK")
		} else {
			fmt.Println("ERROR")
		}

	}

}

func IsCorrectBracket(s string) bool {

	stack := []rune{}

	validate := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		switch char {
		case '(', '{', '[':
			stack = append(stack, char)
		case ')', '}', ']':
			if len(stack) == 0 || stack[len(stack)-1] != validate[char] {
				return false
			}
			stack = stack[:len(stack)-1]

		}
	}
	return len(stack) == 0
}
