package main

import "fmt"
var x int

func init() { // o que manda é o nome init, se coloco func n por exemplo, não funciona e retorna 0 o valor de x
	fmt.Println("função init sendo executada")
	x = 10
}

func main() {
	fmt.Println("função main sendo executada")
	fmt.Println(x)
}

//output
//função init sendo executada
// função main sendo executada
//10
