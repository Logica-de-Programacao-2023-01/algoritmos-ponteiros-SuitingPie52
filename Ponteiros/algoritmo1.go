package main

import "fmt"

// Escreva uma função que receba um ponteiro para um inteiro e um valor inteiro ``n``.
// A função deve atualizar o valor do inteiro para a soma dos ``n`` primeiros números naturais.

func somaNaturais(p *int, n int) {

	soma := 0

	for i := 1; i <= n; i++ {

		soma += i

	}

	*p = soma

}

func main() {

	x := 0

	somaNaturais(&x, 5)

	fmt.Println(x)

}
