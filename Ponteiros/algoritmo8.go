package main

import "fmt"

// Crie uma função que receba um ponteiro para uma variável como argumento
// e modifique o valor da variável dentro da função.
// Certifique-se de que o ponteiro aponte para uma área de memória válida
// e que a memória seja liberada após o uso.

func alterar(v *int) {

	x := 0

	fmt.Println("Digite um valor para a variável:")
	fmt.Scanln(&x)

	*v = x

}

func main() {

	x := 0

	alterar(&x)

	fmt.Println(x)

}
