package main

import (
	"fmt"
	"math"
)

// Escreva uma função em Go que receba um ponteiro para uma variável float64
// e atualize o valor da variável para a média aritmética entre o seu valor atual e o valor da constante Pi.

func mediaPi(n *float64) {

	*n = (math.Pi + *n) / 2

}

func main() {

	x := 34.32

	mediaPi(&x)

	fmt.Println(x)

}
