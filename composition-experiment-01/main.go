package main

import (
	"experiments/composition-experiment-01/models"
	"fmt"
)

func main (){
	b := models.Beverage{Color: "yellow"}
	polar := models.Brand{Brandname: "Polar", Country: "Venezuela"}
	beer := models.Beer{Name: "Solera Azul", Alcohol: 7, Beverage:b, Brand: polar}
	beer.IsBeer()

	fmt.Printf("\n----Type of b variable is: %T",b)
}
