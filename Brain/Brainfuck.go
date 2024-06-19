package main

import (
	"fmt"
	"os"
)

var pointer, progCount int
var pointByte [30000]byte

func Scan(dir int, prog string) {
	nest := 1
	for nest > 0 {
		progCount += dir
		if progCount < 0 || progCount >= len(prog) {
			fmt.Println("Error: unmatched '[' or ']'")
			os.Exit(1)
		}
		switch prog[progCount] {
		case '[':
			if dir == 1 {
				nest++
			} else {
				nest--
			}
		case ']':
			if dir == 1 {
				nest--
			} else {
				nest++
			}
		}
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <brainfuck_code>")
		return
	}
	prog := os.Args[1]

	for progCount < len(prog) {
		switch prog[progCount] {
		case '>':
			pointer++
			if pointer >= len(pointByte) {
				pointer = 0
			}
		case '<':
			if pointer == 0 {
				pointer = len(pointByte) - 1
			} else {
				pointer--
			}
		case '+':
			pointByte[pointer]++
		case '-':
			pointByte[pointer]--
		case '.':
			fmt.Print(string(pointByte[pointer]))
		case ',':
			input := make([]byte, 1)
			if _, err := os.Stdin.Read(input); err != nil {
				fmt.Println("Error reading input:", err)
				return
			}
			pointByte[pointer] = input[0]
		case '[':
			if pointByte[pointer] == 0 {
				Scan(1, prog)
			}
		case ']':
			if pointByte[pointer] != 0 {
				Scan(-1, prog)
			}
		}
		progCount++
	}
}
