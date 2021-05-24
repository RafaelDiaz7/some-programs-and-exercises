package models

import "fmt"

type Beer struct {
	Name string
	Alcohol float32
	//price
	Beverage
	Brand
}

func (b *Beer) IsBeer()  {
	fmt.Println("The", b.Name,"beer")
	fmt.Println("-color:",b.Color)
	fmt.Println("-alcohol grade:",b.Alcohol)
	fmt.Println("-brand:",b.Brandname)
	fmt.Println("-country:",b.Country)
}

