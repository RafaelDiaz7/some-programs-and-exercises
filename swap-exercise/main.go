package main

import (
	"fmt"
)

func swap(xPtr *int,yPtr *int) {
	cache := *xPtr
	//fmt.Println(reflect.TypeOf(cache))
	*xPtr = *yPtr
	*yPtr = cache
}

func main()  {
	x := 1
	y := 2
	swap(&x,&y)
	fmt.Println(x,y)
}
