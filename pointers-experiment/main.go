package main

import "fmt"

func zero(xPtr *int)  {
	*xPtr = 0
}

func four(xPtr *int)  {
	*xPtr = 4
}

func main()  {
	x := 1
	fmt.Println(x)
	zero(&x)
	fmt.Println(x) // Its going to happen that x value would not be changed. X its still 1.
	four(&x)
	fmt.Println(x)
}
