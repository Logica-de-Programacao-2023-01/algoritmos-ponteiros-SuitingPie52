package main

import "fmt"

// Escreva uma função que receba um ponteiro para um inteiro e verifique se esse inteiro é par ou ímpar.
// A função deve atualizar o valor do inteiro para 0 se ele for par ou para 1 se for ímpar.

func oddOrEven(n *int) {

	if *n%2 == 0 {

		*n = 0

	} else {

		*n = 1

	}

}

func main() {

	x := 2
	y := 3

	oddOrEven(&x)

	fmt.Println(x)

	oddOrEven(&y)

	fmt.Println(y)

}
