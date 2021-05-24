package models

import "fmt"

type Beverage struct {
	Color string
}

func (b *Beverage) isColor()  {
	fmt.Println(b.Color)
}
