package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	if len(os.Args) != 2 {
		return
	}

	nb := Atoi(os.Args[1])

	if nb == 0 || nb == 1 {
		return
	}
	var factors []string

	for i := 2; i*i <= nb; i++ {
		for nb%i == 0 {
			factors = append(factors, Itoa(i))
			nb /= i
		}

	}
	if nb > 2 {
		factors = append(factors, Itoa(nb))
	}
	for i, sl := range factors {
		if i != len(factors)-1 {
			Printn(sl)
			Printn("*")
		} else {
			Printn(sl)
		}
	}
	Printn("\n")

}
func Atoi(s string) int {
	num := 0

	for _, v := range s {
		n := int(v - 48)

		if n < 0 || n > 9 {
			return 0
		}
		num = (num * 10) + n

	}
	return num

}
func Itoa(n int) string {
	s := ""
	if n == 0 {
		s = string(rune(48)) + s
	}
	for n > 0 {
		w := n % 10
		s = string(rune(w+48)) + s
		n = n / 10
	}

	return s

}
func Printn(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}

}


// Seth Solution

package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	num, err := strconv.Atoi(os.Args[1])
	if err != nil || num <= 0 {
		return
	}

	factors := primeFactors(num)
	if len(factors) == 0 {
		return
	}

	for i, factor := range factors {
		for _, n := range Itoa(factor) {
			z01.PrintRune(n)
		}
		if i < len(factors)-1 {
			z01.PrintRune('*')
		}
	}
	z01.PrintRune('\n')
}

// Function to find prime factors of a number
func primeFactors(n int) []int {
	factors := make([]int, 0)
	for i := 2; i <= n; i++ {
		for n%i == 0 {
			factors = append(factors, i)
			n /= i
		}
	}
	return factors
}

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
