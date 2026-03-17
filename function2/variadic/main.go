package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func escrever(txt string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(txt, numero)
	}
}

func main() {
	totalDasoma := soma(1, 2, 3, 4, 5, 6, 200, 102, 12, 13)
	fmt.Println(totalDasoma)

	escrever("hello world", 1, 2, 4, 5, 6, 7)
}
