package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 2 || os.Args[1][0] == '-' {
		fmt.Println(0)
		return
	}

	nb := Atoi(os.Args[1])
	if nb == 0 || nb == 1 {
		fmt.Println(0)
		return
	}
	sum := 0

	for i := nb; i > 1; i-- {
		if Isprime(i) {
			sum += i
		}
	}
	fmt.Println(sum)

}

func Isprime(n int) bool {

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true

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

// Seth solution

package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}

	args := os.Args[1]

	num, err := strconv.Atoi(args)
	if err != nil {
		z01.PrintRune('0')
		return
	}

	primeSum := 0
	for i := 2; i <= num; i++ {
		if IsPrime(i) {
			primeSum += i
		}
	}
	for _, c := range Itoa(primeSum) {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}

func IsPrime(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func Itoa(n int) string {
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	} else if n == 0 {
		return "0"
	}

	var digits []rune

	for n > 0 {
		digit := n % 10
		digits = append([]rune{rune('0' + digit)}, digits...)
		n /= 10
	}
	return string(digits)
}

