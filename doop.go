package main

import (
	"os"
	"strconv"
)

//to add Atoi manual
func doop(s []string) {
	if len(s[0]) > 12 {
		return
	}
	num1, _ := strconv.Atoi(s[0])
	operator := s[1]
	num2, _ := strconv.Atoi(s[2])
	result := 0

	if num1 == 0 || num2 == 0 && operator == "/" {
		os.Stdout.Write([]byte("No division by 0\n\n"))
		return
	}
	if num1 == 0 || num2 == 0 && operator == "%" {
		os.Stdout.Write([]byte("No modulo by 0\n\n"))
		return
	}
	switch operator {
	case "+":
		result = num1 + num2
		os.Stdout.Write([]byte(Itoa(result)))
		os.Stdout.Write([]byte("\n"))
	case "-":
		result = num1 - num2
		os.Stdout.Write([]byte(Itoa(result)))
		os.Stdout.Write([]byte("\n"))

	case "*":
		result = num1 * num2
		os.Stdout.Write([]byte(Itoa(result)))
		os.Stdout.Write([]byte("\n"))

	case "/":
		result = num1 / num2
		os.Stdout.Write([]byte(Itoa(result)))
		os.Stdout.Write([]byte("\n"))

	case "%":
		result = num1 % num2

		os.Stdout.Write([]byte(Itoa(result)))
		os.Stdout.Write([]byte("\n"))

	default:
		return
	}
}

func main() {
	if len(os.Args) != 4 {
		return
	}
	doop(os.Args[1:])
}

func Itoa(n int) string {
	//declare variables
	r := []byte{}
	temps := []byte{}
	right := 0

	//handle -ve numbers
	if n < 0 {
		r = append(r, '-')
		n *= -1
	}

	//convert each digit to a string by adding  "0"
	for {
		right = n % 10
		n = n / 10
		temps = append(temps, byte('0'+right))
		if n == 0 {
			break
		}

	}
	for i := len(temps) - 1; i >= 0; i-- {
		r = append(r, temps[i])
	}
	return string(r)
 }

//rays solution
// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {

// 	if len(os.Args[1:]) != 3 {
// 		return
// 	}

// 	arg1 := os.Args[1]
// 	operator := os.Args[2]
// 	arg2 := os.Args[3]
// 	var num1 int
// 	var num2 int
// 	var ans int

// 	for i, v := range arg1 {
// 		if i == 0 && v == '-' {
// 			continue
// 		}
// 		num := int(v - 48)
// 		if num > 9 {
// 			return
// 		}
// 		num1 = (num1 * 10) + num

// 	}

// 	if arg1[0] == '-' {
// 		num1 = -1 * num1
// 	}

// 	for i, v := range arg2 {

// 		if i == 0 && v == '-' {
// 			continue
// 		}

// 		num := int(v - 48)

// 		if num > 9 {
// 			return
// 		}
// 		num2 = (num2 * 10) + num

// 	}
// 	if arg2[0] == '-' {
// 		num2 = -1 * num2
// 	}
// 	//fmt.Println(num1)
// 	//fmt.Println(num2)

// 	if (num1 >= 9223372036854775807 || num1 <= -9223372036854775807) || (num2 >= 9223372036854775807 || num2 <= -9223372036854775807) {
// 		return
// 	}

// 	if (num1 == 0 || num2 == 0) && operator == "/" {
// 		fmt.Println("No division by 0")
// 		return
// 	} else if (num1 == 0 || num2 == 0) && operator == "%" {
// 		fmt.Println("No modulo by 0")
// 		return
// 	}
// 	switch operator {
// 	case "+":
// 		ans = num1 + num2
// 	case "-":
// 		ans = num1 - num2
// 	case "/":
// 		ans = num1 / num2
// 	case "%":
// 		ans = num1 % num2
// 	case "*":
// 		ans = num1 * num2
// 	default:
// 		return
// 	}
// 	fmt.Println(ans)

// }

// Seth Solution

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	value1 := Atoi(os.Args[1])
	operator := os.Args[2]
	value2 := Atoi(os.Args[3])

	if (value1 >= 9223372036854775807 || value1 <= -9223372036854775807) || (value2 >= 9223372036854775807 || value2 <= -9223372036854775807) {
		return
	}

	var result int
	switch operator {
	case "+":
		result = value1 + value2
	case "-":
		result = value1 - value2
	case "*":
		result = value1 * value2
	case "/":
		if value2 == 0 {
			word := "No division by 0"
			for _, c := range word {
				z01.PrintRune(c)
			}
			z01.PrintRune('\n')
			return
		}
		result = value1 / value2
	case "%":
		if value2 == 0 {
			word := "No modulo by 0"
			for _, c := range word {
				z01.PrintRune(c)
			}
			z01.PrintRune('\n')
			return
		}
		result = value1 % value2
	default:
		return
	}
	r := Itoa(int(result))
	for _, c := range r {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}

func Itoa(n int) string {
	if n == 0 {
		return "0"
	} else if n < 0 {
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

func Atoi(s string) int {
	var number int
	sign := 1

	for idx, char := range s {
		if char == '-' && idx == 0 {
			sign = -1
		} else if char == '+' && idx == 0 {
			sign = 1
		} else if char >= '0' && char <= '9' {
			number = number*10 + int(char-'0')
		} else {
			return 0
		}
	}
	return number * sign
}

//Stella's Soln 

package main

import (
	"os"
)

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}
	sign := ""
	if n < 0 {
		sign = "-"
		n = -n
	}
	q := ""
	for n > 0 {
		digits := n % 10
		q = string(rune('0'+digits)) + q
		n /= 10
	}
	return sign + q
}

func Atoi(s string) (int, bool) {
	sign := 1
	q := 0

	for i, v := range s {
		if v == '-' && i == 0 {
			sign = -1
		} else if v == '+' && i == 0 {
			sign = 1
		} else if v >= '0' && v <= '9' {
			q = q*10 + int(v-'0')
		} else {
			return 0, false
		}
	}
	return sign * q, true
}

func PrintStr(s string) {
	os.Stdout.WriteString(s)
	os.Stdout.WriteString("\n")
}

func Calculation(a, operator, b string) string {
	value1, ok1 := Atoi(a)
	value2, ok2 := Atoi(b)

	if !ok1 || !ok2 {
		return ""
	}

	// Define max int64 value to handle large numbers
	const maxInt64 = 9223372036854775807

	if (value1 >= maxInt64 || value1 <= -maxInt64) || (value2 >= maxInt64 || value2 <= -maxInt64) {
		return ""
	}

	q := 0
	switch operator {
	case "+":
		q = value1 + value2
	case "-":
		q = value1 - value2
	case "*":
		q = value1 * value2
	case "/":
		if value2 == 0 {
			return "No division by 0"
		}
		q = value1 / value2
	case "%":
		if value2 == 0 {
			return "No modulo by 0"
		}
		q = value1 % value2
	default:
		return ""
	}
	return Itoa(q)
}

func main() {
	if len(os.Args) != 4 {
		return
	}
	value1, operator, value2 := os.Args[1], os.Args[2], os.Args[3]

	result := Calculation(value1, operator, value2)

	if result != "" {
		PrintStr(result)
	}
}

