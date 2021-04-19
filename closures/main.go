package main

import "fmt"

func main()  {
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven()) //0
	fmt.Println(nextEven()) //2
	fmt.Println(nextEven()) //4
	fmt.Println(nextEven()) //6
	fmt.Println(nextEven()) //6
	fmt.Println()
	nextOdd:= makeOddGenerator()
	fmt.Println(nextOdd()) //0
	fmt.Println(nextOdd()) //1
	fmt.Println(nextOdd()) //3
	fmt.Println(nextOdd()) //5
	fmt.Println(nextOdd()) //7
	fmt.Println(nextOdd()) //7

}

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func makeOddGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 1
		if i != 0 && i != 1 {
			i+=1
		}
		return
	}
}