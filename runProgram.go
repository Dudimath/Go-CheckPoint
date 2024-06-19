package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	file, err := os.Open("m.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	//Check if the directory exists
	//if not create it
	if _, err := os.Stat("test"); os.IsNotExist(err) {
		if err := os.Mkdir("test", 0777); err != nil {
			panic(err)
		}

	}
	//change directory to the created directory
	if err := os.Chdir("test"); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	function, err := os.Create("main.go")
	if err != nil {
		panic(err)
	}
	defer function.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//write every read line in its own line
		function.WriteString(scanner.Text() + "\n")
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	//cmd := exec.Command("gofmt", "-d", "-w", "../test")
	//execute the go run command with the name of the created go file
	cmd := exec.Command("go", "run", "main.go")
	output, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
	}
	if len(output) < 1 {
		fmt.Println("Well formatted!")
	}
	// Print the output
	fmt.Printf("%s", output)

}
