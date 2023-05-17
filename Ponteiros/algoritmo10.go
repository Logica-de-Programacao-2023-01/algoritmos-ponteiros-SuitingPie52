package main

import "fmt"

// Implemente uma função que receba um ponteiro para uma slice
// e seu tamanho e preencha-o com os ``n`` primeiros números primos.

func primos(s *[]int, n int) {

	primo := true

	for i := 2; len(*s) != n; i++ {

		for j := 2; j < i; j++ {

			if i%j == 0 {

				primo = false
				break

			}

		}

		if primo == true {

			*s = append(*s, i)

		} else {

			primo = true

		}

	}

}

func main() {

	conj := []int{}

	primos(&conj, 10)

	fmt.Println(conj)

}
