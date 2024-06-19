
// Instructions

//     Write a function that simulates the behavior of the Itoa function in Go. Itoa transforms a number represented as anint in a number represented as a string.

//     For this exercise the handling of the signs + or - does have to be taken into account.

// Expected function

// func Itoa(n int) string {

// }

// Usage

// Here is a possible program to test your function :

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
//     fmt.Println(piscine.Itoa(12345))
//     fmt.Println(piscine.Itoa(0))
//     fmt.Println(piscine.Itoa(-1234))
//     fmt.Println(piscine.Itoa(987654321))
// }

// And its output :

// $ go run .
// 12345
// 0
// -1234
// 987654321
// $

package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}

	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}

	var digits []rune
	for n > 0 {
		digit := n % 10
		digits = append([]rune{rune('0' + digit)}, digits...)
		n /= 10
	}
	return string(digits)
}

func main() {
	fmt.Println(Itoa(12345))
	fmt.Println(Itoa(0))
	fmt.Println(Itoa(-1234))
	fmt.Println(Itoa(987654321))
}
