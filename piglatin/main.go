package main

import (
	"fmt"
	"os"
)

func pigLatin(s string) (result string) {
	// handle wordswith no vowels
	check := false
	for _, char := range s {
		if checkVowel(char) {
			check = true

		}
		if check {
			break
		}

	}
	if !check {
		result = ("No vowels")
		return result
	}
	//handle words that start with vowels
	if checkVowel(rune(s[0])) {
		result = s + "ay"
		return result
	}
	//handle words that have vowels in the middle
	for i, char := range s {
		if i != 0 && checkVowel(char) {
			v := s[i:]
			s = s[:i]
			result = v + s + "ay"
		}
	}

	return result
}

//function to check for vowels
func checkVowel(s rune) bool {
	return s == 'a' || s == 'e' || s == 'i' || s == 'o' || s == 'u'
}

func main() {
	if len(os.Args) != 2 {
		return
	}
	args := os.Args[1]
	fmt.Println(pigLatin(args))
}

// rays solution
// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {

// 	if len(os.Args) != 2 {
// 		return
// 	}

// 	arg := os.Args[1]

// 	fmt.Println(PigLatin(arg))

// }
// func CheckVowel(a rune) bool {
// 	return a == 'a' || a == 'e' || a == 'i' || a == '0' || a == 'u' || a == 'A' || a == 'E' || a == 'I' || a == 'O' || a == 'U'
// }

// func PigLatin(s string) string {
// 	var result string
// 	var w string

// 	for i, ch := range s {

// 		if i == 0 && CheckVowel(ch) {
// 			result = s + "ay"
// 			break

// 		} else if i > 0 && CheckVowel(ch) {
// 			w = s[i:]
// 			s = s[:i]
// 			result = w + s + "ay"
// 			break
// 		} else {
// 			result = "No vowel"
// 		}

// 	}
// 	return result

// }

// Seth Solution

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	args := os.Args[1]

	word := ""
	for i, c := range args {
		if checkVowels(c) && i == 0 {
			word = args + "ay"
		} else if checkVowels(c) && i != 0 {
			word = args[i:] + args[:i] + "ay"
		}
	}

	check := false
	for _, c := range word {
		if checkVowels(c) {
			check = true
			break
		}
	}

	if !check {
		word = "No vowels"
	}

	for _, c := range word {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}

func checkVowels(c rune) bool {
	return (c == 'a' || c == 'A' || c == 'e' || c == 'E' || c == 'i' || c == 'I' || c == 'o' || c == 'O' || c == 'u' || c == 'U')
}

