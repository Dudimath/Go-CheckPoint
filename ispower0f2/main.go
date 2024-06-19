package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	arg := os.Args[1]
	newArg := 0
	for _, v := range arg {
		n := int(v - 48)

		newArg = (newArg * 10) + n

	}

	fmt.Println(isPowerof2(newArg))
}

func isPowerof2(newArg int) bool {
	return newArg%2 == 0
}

// Seth  Solution
// package main

// import (
// 	"os"

// 	"github.com/01-edu/z01"
// )

// func main() {
// 	if len(os.Args) != 2 {
// 		return
// 	}
// 	args := os.Args[1]
// 	if Atoi(args) <= 0 {
// 		return
// 	}

// 	word := ""
// 	if IsPowerOf2(Atoi(args)) {
// 		word = "true"
// 	} else {
// 		word = "false"
// 	}

// 	for _, c := range word {
// 		z01.PrintRune(c)
// 	}
// 	z01.PrintRune('\n')
// }

func IsPowerOf2(num int) bool {
	return num > 0 && (num&(num-1)) == 0
}

// func Atoi(s string) int {
// 	var number int
// 	sign := 1

// 	for idx, char := range s {
// 		if char == '-' && idx == 0 {
// 			sign = -1
// 		} else if char == '+' && idx == 0 {
// 			sign = 1
// 		} else if char >= '0' && char <= '9' {
// 			number = number * 10 + int(char -'0')
// 		} else {
// 			return 0
// 		}
// 	}
// 	return number * sign
// }

// ombima solution

// package main

// import (
// "os"

// "github.com/01-edu/z01"
// )

func main() {
	if len(os.Args) != 2 {
		return
	}

	args := os.Args[1]

	n := BasicAtoi(args)
	if isPowerOf2(n) {
		printStr("true")
	} else {
		printStr("false")
	}
}

func isPowerOf2(n int) bool {
	if n <= 0 {
		return false
	}
	return (n & (n - 1)) == 0
}

func BasicAtoi(s string) int {
	var number int
	for _, ch := range s {
		number = number*10 + int(ch-'0')
	}
	return number
}

func printStr(s string) {
	for _, ch := range s {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
