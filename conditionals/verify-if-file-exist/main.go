package main

import (
	"fmt"
	"os"
)

func main() {
	var fileRoute string
	fileRoute = "main.c"

	if _, err := os.Stat(fileRoute); os.IsNotExist(err) {
		fmt.Println(fileRoute, "not exist.")
	} else {
		fmt.Println(fileRoute, "exist.")
	}
}
