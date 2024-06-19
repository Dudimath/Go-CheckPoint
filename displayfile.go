// displayfile
// Instructions

// Write a program that displays, on the standard output, the content of a file given as argument.
// Usage :

// $ go run .
// File name missing
// $ echo "Almost there!!" > quest8.txt  // USE echo 'Almost there!!' > quest8.txt    : bash terminal treats !! as a coomand to include previously run commands in the file being echoed to when you use ""
// $ go run . quest8.txt main.go
// Too many arguments
// $ go run . quest8.txt
// Almost there!!

package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("File name missing")
		return
	}

	if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
	}

	fileContent, err := os.ReadFile("quest8.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(fileContent))
}
