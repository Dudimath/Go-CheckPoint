package main

import "fmt"

func main() {
	octet := byte(0b01000001) // Example byte
	swapped := SwapBits(octet)
	fmt.Printf("Original byte: %08b\n", octet)
	fmt.Printf("Swapped byte: %08b\n", swapped)
}

func SwapBits(octet byte) byte {

	upperHalf := octet >> 4
	lowerHalf := octet & 0b00001111

	swappedBits := (lowerHalf << 4) | upperHalf

	return swappedBits

}

// Seth Solution

func SwapBits(octet byte) byte {
	return (octet << 4) | (octet >> 4)
}
func main() {} //include an empty function main before submitting your code.
