package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		return
	}
	args := os.Args[1:]
	s := ""
	for _, arg := range args {
		s = s + arg
		s += " "

	}
	fmt.Println(rotateVowel(s))
}
func rotateVowel(s string) string {
	sb := []byte(s)

	n := len(sb) - 1
	mid := 0
	if n%2 == 0 {
		mid = n / 2
	} else {
		mid = n/2 + 1
	}

	right := n

	count := 0
	for i := 0; i <= mid; i++ {
		count++
		if checkVowels(sb[i]) {
			for j := right; j >= mid; j-- {
				count++
				if checkVowels(sb[j]) {
					sb[i], sb[j] = sb[j], sb[i]
					right = j - 1
					break
				}
			}
		}
	}
	return string(sb)
}
func checkVowels(v byte) bool {
	return v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' || v == 'A' || v == 'E' || v == 'I' || v == 'O' || v == 'U'
}
