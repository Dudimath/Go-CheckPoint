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

	chars := make(map[rune]bool)
	ans := make([]rune, 0, len(str1))

	for _, ch := range str1 {

		if _, oK := chars[ch]; !oK && containsRune(str2, ch) {
			chars[ch] = true
			ans = append(ans, ch)
		}
	}
	fmt.Println(string(ans))

}

func containsRune(s string, r rune) bool {
	for _, char := range s {
		if char == r {
			return true
		}
	}
	return false
}

// Seth Solution

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	args1 := os.Args[1]
	args2 := os.Args[2]
	seen := make(map[rune]bool)
	exist := make(map[rune]bool)
		
	for _, char := range args2 {
		exist[char] = true
	}

	for _, char := range args1 {
		if exist[char] && !seen[char] {
			seen[char] = true
			z01.PrintRune(char)
		}
	}
	z01.PrintRune('\n')
}

//Stella's Soln

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	args1, args2, seen, exist := os.Args[1], os.Args[2], make(map[rune]bool), make(map[rune]bool)

	for _, char := range args2 {
		exist[char] = true
	}

	for _, char := range args1 {
		if exist[char] && !seen[char] {
			seen[char] = true
			z01.PrintRune(char)
		}
	}
	z01.PrintRune('\n')
}

