// foldint
// Instructions

// The function should have as parameters a function, f func(int, int) int a slice of integers, slice []int and an int acc int. For each element of the slice, it should apply the arithmetic function, save the result and print it. The function will be tested with our own functions Add, Sub, and Mul.
// Expected function

// func FoldInt(f func(int, int) int, a []int, n int) {

// }

// Usage

// Here is a possible program to test your function:

// package main

// func main() {
// 	table := []int{1, 2, 3}
// 	ac := 93
// 	FoldInt(Add, table, ac)
// 	FoldInt(Mul, table, ac)
// 	FoldInt(Sub, table, ac)
// 	fmt.Println()

// 	table = []int{0}
// 	FoldInt(Add, table, ac)
// 	FoldInt(Mul, table, ac)
// 	FoldInt(Sub, table, ac)
// }

// And its output :

// $ go run .
// 99
// 558
// 87

// 93
// 0
// 93
// $

// Ray Soluton

package main

import (
	"fmt"
)

func main() {
	table := []int{1, 2, 3}
	ac := 93
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
	fmt.Println()

	table = []int{0}
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
}

func FoldInt(f func(int, int) int, a []int, n int) {
	for _, v := range a {
		n = f(n, v)
	}
	fmt.Println(n)
}

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func Mul(a, b int) int {
	return a * b
}

// Seth Solution

package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func FoldInt(f func(int, int) int, a []int, n int) {
	var tot int
	for _, y := range a {
		tot += y
	}
	result := f(tot, n)
	for _, c := range Itoa(result) {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}

func Add(a, b int) int {
	return a + b
}

func Mul(a, b int) int {
	return a * b
}

func Sub(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func Itoa(n int) string {
	if n == 0 {
		return "0"
	} else if n < 0 {
		z01.PrintRune('-')
		n = -n
	}

	var digits []rune 
	for n > 0 {
		digit := n%10
		digits = append([]rune{rune('0' + digit)}, digits...)
		n /= 10
	}
	return string(digits)
}

func main() {
	table := []int{1, 2, 3}
	ac := 93
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
	fmt.Println()

	table = []int{0}
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
}

// Vincent Solution(short code)
package main

import "fmt"

func FoldInt(f func(int, int) int, a []int, n int) {
	ac := n
	for _, v := range a {
		ac = f(ac, v)
	}
	fmt.Println(ac)
}

func Add(ac int, v int) int { return ac + v }
func Sub(ac int, v int) int { return ac - v }
func Mul(ac int, v int) int { return ac * v }
func Div(ac int, v int) int { return ac / v }

func main() {
	table := []int{1, 2, 3}
	ac := -93
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
	fmt.Println()

	table = []int{0}
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
}

