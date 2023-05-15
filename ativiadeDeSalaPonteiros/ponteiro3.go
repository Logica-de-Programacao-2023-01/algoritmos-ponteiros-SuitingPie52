package main

import "fmt"

type Produto struct {
	nome    string
	preco   float64
	estoque int
}

func trocarPreco(p *Produto, valor float64) {

	p.preco = valor

}

func main() {

	chinelo := Produto{

		nome:    "Chinelo",
		preco:   20,
		estoque: 2,
	}

	trocarPreco(&chinelo, 100)

	fmt.Println(chinelo.preco)

}