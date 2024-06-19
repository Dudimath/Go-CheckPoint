package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	if len(os.Args) != 2 {
		return
	}

	str := os.Args[1]
	ans := ""
	if str == "" {
		z01.PrintRune('\n')
		return
	}
	for _, ch := range str {
		if Alpha(ch) && Capital(ch) {
			rpt := int(ch - 'A')
			for i := 0; i <= rpt; i++ {
				ans += string(ch)
			}
		} else if Alpha(ch) && !Capital(ch) {
			rpt := int(ch - 'a')
			for i := 0; i <= rpt; i++ {
				ans += string(ch)
			}
		} else {
			ans += string(ch)
		}

	}
	Printn(ans)

}
func Alpha(c rune) bool {
	return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z')
}
func Capital(c rune) bool {
	return (c >= 'A' && c <= 'Z')
}

func Printn(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
