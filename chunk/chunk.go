// chunk
// Instructions
// Write a function called Chunk that receives as parameters a slice, slice []int, and a number size int. The goal of this function is to chunk a slice into many sub slices where each sub slice has the length of size.

// If the size is 0 it should print a newline ('\n').
// Expected function
// func Chunk(slice []int, size int) {

// }
// Usage
// Here is a possible program to test your function :

// package main

// func main() {
// 	Chunk([]int{}, 10)
// 	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
// 	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
// 	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
// 	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
// }
// And its output :

// $ go run .
// []

// [[0 1 2] [3 4 5] [6 7]]
// [[0 1 2 3 4] [5 6 7]]
// [[0 1 2 3] [4 5 6 7]]
// $

// Bravian Solution

package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func Chunk(slice []int, size int) {
	if size == 0 {
		fmt.Println()
		return
	}
	//ceiling algorithm to get the size of the slice of slices
	m := len(slice) / size
	rem := len(slice) % size
	if rem >= 1 {
		m++
	}
//initializing the slice of slices
	maped := make([][]int, m)

	n := 0//the index of the slice currently being worked on
	count := 0//to compare to size

	for i := 0; i < len(slice); i++ {
		if count == size {
			n++
			count = 0
			i--
		} else {
			maped[n] = append(maped[n], slice[i])
			count++
		}
	}
	println(maped)
}

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}

// seth solution

package main

import (
	"github.com/01-edu/z01"
)

func Chunk(slice []int, size int) {
	if size == 0 {
		z01.PrintRune('\n')
		return
	}

	var chunkSlice [][]int

	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		chunkSlice = append(chunkSlice, slice[i:end])
	}

	z01.PrintRune('[')
	for j, c := range chunkSlice {
		z01.PrintRune('[')
		for i, v := range c {
			for _, r := range Itoa(v) {
				z01.PrintRune(rune(r))
			}
			if i < len(c)-1 {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune(']')
		if j < len(chunkSlice)-1 {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune(']')
	z01.PrintRune('\n')
}

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
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

//Rays solution

func Chunk(slice []int, size int) {
	count := 0
	var newSlice []int
	var outputSlice [][]int
	if size == 0 {
		fmt.Println()
		return
	}
	if len(slice) == 0 {
		fmt.Println(slice)
		return

	}

	for _, n := range slice {
		newSlice = append(newSlice, n)
		count++
		if count == size {
			outputSlice = append(outputSlice, newSlice)
			newSlice = []int{}
			count = 0
		}

	}
	if len(newSlice) > 0 {
		outputSlice = append(outputSlice, newSlice)
	}

	fmt.Println(outputSlice)

}

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}
