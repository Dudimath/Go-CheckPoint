package main

import (
	"github.com/01-edu/z01"
)

func main() {
	//fmt.Println(ConveToHexadecimal("l"))
	// /PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})

}

func ConveToHexadecimal(s string) string {
	hexadec := "0123456789abcdef"

	hexdec := ""

	if s == "" {
		return "00"
	}

	for _, r := range s {

		for r > 0 {
			index := r % 16

			hexdec = string(hexadec[index]) + hexdec

			r /= 16

		}

	}

	return hexdec
}

func PrintMemory(arr [10]byte) {

	for i, b := range arr {
		// byte 0
		if b == '\x00' {
			Printn("00")
		}

		//fmt.Printf("%02x", b)
		Printn(ConveToHexadecimal(string(b)))

		if (i+1)%4 == 0 {
			z01.PrintRune('\n')
		} else if i != len(arr)-1 {
			z01.PrintRune(' ')
		}

	}
	Printn("\n")

	for _, b := range arr {
		if b >= 32 && b <= 126 {
			Printn(string(b))
		} else {
			Printn(".")
		}
	}
	Printn("\n")

}

func Printn(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
	//z01.PrintRune('\n')
}
