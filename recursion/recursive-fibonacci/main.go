package main

import "fmt"

/*
Introducing go exercise pg 53
5. The Fibonacci sequence is defined as: 
fib(0) = 0, fib(1) = 1, fib(n) = fib(n-1) + fib(n-2).
Write a recursive function that can find fib(n). 

*/

func main() {
	fmt.Println(fib(10))
}

// This recursive fibonacci its working well! I will also make the dynamic fibonacci luegoo
func fib(n uint32) uint32 {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}
