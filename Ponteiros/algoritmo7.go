package main

import "fmt"

// Escreva uma função em Go que receba um ponteiro para um struct Conta com campos saldo e titular,
// e adicione um valor ao saldo da conta.

type Conta struct {
	Saldo   float64
	Titular string
}

func addValor(c *Conta) {

	x := 0.0

	fmt.Println("Digite um valor para adicionar a sua conta:")
	fmt.Scanln(&x)

	c.Saldo += x

}

func main() {

	c := Conta{
		Saldo:   231,
		Titular: "Gabriel",
	}

	addValor(&c)

	fmt.Println(c.Saldo)

}
