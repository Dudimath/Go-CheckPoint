package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 3 {
		return
	}

	str1 := os.Args[1]
	str2 := os.Args[2]

	ptr1, ptr2 := 0, 0

	for ptr2 < len(str2) {

		if ptr1 < len(str1) && str1[ptr1] == str2[ptr2] {
			ptr1++
		}
		ptr2++
	}

	if ptr1 == len(str1) {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}

}

// Seth solution

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	arg1, arg2 := os.Args[1], os.Args[2]

	i, j := 0, 0

	word := ""
	for i < len(arg1) && j < len(arg2) {
		if arg1[i] == arg2[j] {
			word += string(arg1[i])
			i++
		}
		j++
	}

	if arg1 == word {
		z01.PrintRune('1')
	} else {
		z01.PrintRune('0')
	}
	z01.PrintRune('\n')
}
