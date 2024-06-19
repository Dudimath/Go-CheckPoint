// findnextprime
// Instructions

// Write a function that returns the first prime number that is equal or superior to the int passed as parameter.

// The function must be optimized in order to avoid time-outs with the tester.

// (We consider that only positive numbers can be prime numbers)
// Expected function

// func FindNextPrime(nb int) int {

// }

// Usage

// Here is a possible program to test your function :

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	fmt.Println(piscine.FindNextPrime(5))
// 	fmt.Println(piscine.FindNextPrime(4))
// }

// And its output :

// $ go run .
// 5
// 5
// $

package main

import (
	"fmt"
)

func FindNextPrime(nb int) int {
	if nb < 2 {
		return 0
	}

	for !IsPrime(nb) {
		nb++
	}
	return nb
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

func main() {
	fmt.Println(FindNextPrime(5))
	fmt.Println(FindNextPrime(4))

}
