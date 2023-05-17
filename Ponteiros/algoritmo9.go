package main

import "fmt"

// Implemente uma função que receba um ponteiro para uma struct "Livro" com campos "Título", "Autor" e "Preço",
// e que adicione um desconto de 10% no preço do livro.

type Livro2 struct {
	Titulo string
	Autor  string
	Preco  float64
}

func desconto(l *Livro2) {

	l.Preco -= l.Preco * 0.1

}

func main() {

	l := Livro2{
		Titulo: "dawdawfe",
		Autor:  "davasdvasvvs",
		Preco:  100,
	}

	desconto(&l)

	fmt.Println(l.Preco)

}
