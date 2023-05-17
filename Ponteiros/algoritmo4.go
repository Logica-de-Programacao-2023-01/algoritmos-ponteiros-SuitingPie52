package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Escreva uma função em Go que receba um ponteiro para uma variável inteira
// e atualize o valor da variável para a soma dos valores dos seus dois últimos dígitos
// (por exemplo, se o valor inicial da variável for 1234, o novo valor será 3+4=7).

func somaDoisUltimosAlgarismos(n *int) {

	nString := strconv.Itoa(*n)
	nStringAlgarismos := strings.Split(nString, "")
	a1 := 0
	a2 := 0
	soma := 0

	i := len(nStringAlgarismos) - 1

	a1, _ = strconv.Atoi(nStringAlgarismos[i])
	a2, _ = strconv.Atoi(nStringAlgarismos[i-1])
	soma = a1 + a2

	*n = soma

}

func main() {

	x := 21324123421389

	somaDoisUltimosAlgarismos(&x)

	fmt.Println(x)

}
