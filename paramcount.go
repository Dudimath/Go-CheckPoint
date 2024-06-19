// paramcount
// Instructions
// Write a program that displays the number of arguments passed to it. This number will be followed by a newline ('\n').

// If there is no argument, the program displays 0 followed by a newline.

// Usage
// $ go run . 1 2 3 5 7 24
// 6
// $ go run . 6 12 24 | cat -e
// 3$
// $ go run . | cat -e
// 0$
// $

package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println(0)
		return
	}
	args := os.Args[1:]
	count := rune(len(args)) + '0'
	z01.PrintRune(count)
	z01.PrintRune('\n')
}

/*
Solution 2:
Uses a defined Itoa function instead of strconv package. This enables the program to print and display arguments that are more 9

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Itoa(num int) string {
	if num == 0 {
		return "0"
	}

	isNegativeNum := false

	if num < 0 {
		isNegativeNum = true
		num = -num
	}

	var result []byte

	for num > 0 {
		digit := num % 10
		result = append([]byte{byte('0' + digit)}, result...)
		num /= 10
	}

	if isNegativeNum {
		result = append([]byte{'-'}, result...)
	}
	return string(result)
}

func paramCount() {
	var count int

	arguments := os.Args[1:]

	for range arguments {
		count++
	}

	for _, val := range Itoa(count) {
		z01.PrintRune(val)
	}
	z01.PrintRune('\n')
}

func main() {
	paramCount()
}
*/

//Stella's Soln
package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	count := len(os.Args[1:])
	Paramcount := Itoa(count)
	PrintStr(Paramcount)
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}
	sign := ""
	if n < 0 {
		n = -n
		sign = "-"
	}
	q := ""
	for n > 0 {
		digits := n % 10
		q = string(rune('0'+digits)) + q
		n /= 10
	}
	return sign + q
}
