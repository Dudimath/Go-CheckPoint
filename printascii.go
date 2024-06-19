// print-ascii
// Instructions
// Write a program that prints the ASCII value of a letter passed in the command line

// If the argument is not a letter nothing will be printed
// if the number of arguments is not 1 then nothing will be printed
// Usage
// $ go run .
// $ go run . a
// 97
// $ go run . 'A'
// 65
// $ go run . 'z'
// 122
// $ go run . Z
// 90
// $ go run . 1
// $ go run . "Hello" "Word"

package main

import (
	"os"
	"strconv"
	"github.com/01-edu/z01"
)

func isletter(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func main() {
	if len(os.Args) != 2 {
		return
	}

	arg := os.Args[1]

	if len(arg) != 1 || !isletter(arg[0]) {
		return
	}

	for _, n := range strconv.Itoa(int(arg[0])) {
		z01.PrintRune(n)
	}
	z01.PrintRune('\n')
}
