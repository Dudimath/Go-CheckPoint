// strlen
// Instructions
// Write a function that counts the runes of a string and that returns that count.
// Expected function
// func StrLen(s string) int {

// }
// Usage
// Here is a possible program to test your function :

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	l := piscine.StrLen("Hello World!")
// 	fmt.Println(l)
// }
// And its output :

// $ go run .
// 12
// $

package main

import (
	"fmt"
)

func StrLen(s string) int {
	return len(s)
}

func main() {
	l := StrLen("Hello World!")
	fmt.Println(l)
}

// when using os package
// package main

// import (
// 	"fmt"
// 	"os"
// )

// func StrLen(s string) int {
// 	var count int

// 	for range s {
// 		count++
// 	}
// 	return count
// }

// func main() {
// 	if len(os.Args) <= 1 {
// 		return
// 	}

// 	args := os.Args[1]
// 	fmt.Println(StrLen(args))
// }
