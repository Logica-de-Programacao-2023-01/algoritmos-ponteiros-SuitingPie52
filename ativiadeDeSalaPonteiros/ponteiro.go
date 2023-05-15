package main

import "fmt"

func swap(p1 *int, p2 *int) {

	trocar := *p1
	*p1 = *p2
	*p2 = trocar

}

func main() {

	x := 0
	y := 0

	fmt.Println("Informe uma variável:")
	fmt.Scanln(&x)

	fmt.Println("Informe outra variável:")
	fmt.Scanln(&y)

	swap(&x, &y)

	fmt.Println("Suas variáveis com ordem invertido são: ", x, "e", y)

}
