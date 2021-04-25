package main

import "fmt"

func main() {
	var a uint8 = 0xF
	// I want to add a number to 'a' uint8 variable using binary operators and hexadecimal numbers
	a |= 0x10

	/* The |= operator in this case means: a = a | 0x10. That statement means: assign to 'a' the bits that belong to either 'a' or 0x10 */ 

	//This would print 31 if the operation was succesfull
	fmt.Println(a)
}
