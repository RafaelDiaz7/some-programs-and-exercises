package main

import(
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Ingresa una oracion o una cadena de caracteres: ")
	scanner.Scan()

	input := scanner.Text()
	fmt.Println()
	fmt.Printf("Ingresaste: %s y tiene %d caracteres. \n\n", input, len(input))
	
	fmt.Printf("El tipo de dato de la entrada ya procesada es: %T \n",input)
	fmt.Println("Por lo tanto es una cadena de caracteres.")
}
