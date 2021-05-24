package main

import (
	"fmt"
)

func main()  {
	var a int8 = 0x24
	var b int8 = 0xF
	if a > b{
		fmt.Printf("%d is greater than %d\n\n",a,b)
	}
}
