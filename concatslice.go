// concatslice
// Instructions
// Write a function ConcatSlice() that takes two slices of integers as arguments and returns the concatenation of the two slices.

// Expected function
// func ConcatSlice(slice1, slice2 []int) []int {

// }
// Usage
// Here is a possible program to test your function:

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	fmt.Println(piscine.ConcatSlice([]int{1, 2, 3}, []int{4, 5, 6}))
// 	fmt.Println(piscine.ConcatSlice([]int{}, []int{4, 5, 6, 7, 8, 9}))
// 	fmt.Println(piscine.ConcatSlice([]int{1, 2, 3}, []int{}))
// }
// And its output:

// $ go run .
// [1 2 3 4 5 6]
// [4 5 6 7 8 9]
// [1 2 3]

package main

import "fmt"

func ConcatSlice(slice1, slice2 []int) []int {
	var slice []int
	if len(slice1) > 0 && len(slice2) > 0 {
		slice = append(slice, slice1[:]...)
		slice = append(slice, slice2[:]...)
	} else if len(slice1) > 0 && len(slice2) == 0 {
		slice = append(slice, slice1[:]...)
	} else if len(slice2) > 0 && len(slice1) == 0 {
		slice = append(slice, slice2[:]...)
	}
	return slice
}

func main() {
	fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatSlice([]int{}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{}))
	fmt.Println(ConcatSlice([]int{}, []int{}))
}

//Stella's soln

package main

import "fmt"

func ConcatSlice(slice1, slice2 []int) []int {
    return append(slice1, slice2...)
}

func main() {
	fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatSlice([]int{}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{}))
	fmt.Println(ConcatSlice([]int{}, []int{}))
} 
