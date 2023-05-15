package main

import "fmt"

func trocarBool(p *bool) {

	if *p == true {

		*p = false

	} else {

		*p = true

	}

}

func main() {

	t := true
	f := false

	trocarBool(&t)

	fmt.Println(t)

	trocarBool(&f)

	fmt.Println(f)

}
