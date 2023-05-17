package main

import (
	"fmt"
	"strings"
)

// Escreva uma função em Go que receba um ponteiro para uma string
// e atualize o valor da string para seu reverso.

func inverterString(s *string) {

	newS := strings.Split(*s, "")
	*s = ""

	for i := len(newS) - 1; i >= 0; i-- {

		*s += newS[i]

	}

}

func main() {

	s := "Hello World!"

	inverterString(&s)

	fmt.Println(s)

}
