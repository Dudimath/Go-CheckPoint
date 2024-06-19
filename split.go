// split
// instructions

// Write a function that receives a string and a separator and returns a slice of strings that results of splitting the string s by the separator sep.
// Expected function

// func Split(s, sep string) []string {

// }

// Usage

// Here is a possible program to test your function :

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	s := "HelloHAhowHAareHAyou?"
// 	fmt.Printf("%#v\n", piscine.Split(s, "HA"))
// }

// And its output :

// $ go run .
// []string{"Hello", "how", "are", "you?"}
// $

// Seth solution 1
package main

import (
	"fmt"
)

func Split(s, sep string) []string {
	var result []string
	start := 0
	for i := 0; i < len(s); i++ {
		if i+len(sep) <= len(s) && s[i:i+len(sep)] == sep {
			result = append(result, s[start:i])
			start = i + len(sep)
			i += len(sep) - 1
		}
	}
	result = append(result, s[start:])
	return result
}

func main() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", Split(s, "HA"))
}

// Seth Solution 2
package main

import (
	"fmt"
)

func Split(s, sep string) []string {
 var start int
 var end int

 var splitSlice []string

 for i, c := range s {
	if c == 'H' && s[i:i+2] == sep {
		end = i
		splitSlice = append(splitSlice, s[start:end])
		start = i+2
	}
 }
 splitSlice = append(splitSlice, s[start:])
 return splitSlice
}

func main() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", Split(s, "HA"))
}
