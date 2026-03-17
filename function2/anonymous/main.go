package main

import (
	"fmt"
)

func main() {

	retorno := func(txt string) string {
		return fmt.Sprintf("recebido -> %s %d", txt,12)
	}("passando parâmetro")

	fmt.Println(retorno)
}
