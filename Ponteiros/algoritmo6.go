package main

import "fmt"

// Escreva uma função em Go que receba um ponteiro para um struct Livro com campos título e autor,
// e altere o título do livro para "Desconhecido" se o autor for "Anônimo".

type Livro struct {
	Titulo string
	Autor  string
}

func anonimo(l *Livro) {

	if l.Autor == "Anônimo" {

		l.Titulo = "Desconhecido"

	}

}

func main() {

	l := Livro{

		Titulo: "100 anos de solidão",
		Autor:  "Gabriel García Marquez",
	}

	l2 := Livro{
		Titulo: "ryha9gj90krg",
		Autor:  "Anônimo",
	}

	anonimo(&l)
	anonimo(&l2)

	fmt.Println(l)
	fmt.Println(l2)

}