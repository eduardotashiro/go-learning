package main

import "fmt"

func somar(n1 int8, n2 int8) int8 { // parecido com ts, onde o tipo do retorno eu insiro depois do fechar parenteses do parametro
	return n1 + n2
}

// desde que ambos parametros sejam do mesmo tipo, posso declara assim, tipando apenas o ultimo parametro
func calculosMatematicos(n1, n2 int8) (int8, int8) { //func que retorna dois valores ! a ordem importa nesse caso
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := f("texto da função")
	println(resultado)

	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 15) // como eu disse lá em cima, a ordem importa!! , tanto no momento de inferir a função das nas variaveis, quantos no momento de ''''printa-las'''''
	fmt.Println(resultadoSoma, resultadoSubtracao)                   //printando com base na ordem do retorno da função! (ORDEM DOS FATORES ALTERAM O PRODUTO kk)

	resultadoSoma2, _ := calculosMatematicos(10, 15) //até que enfim descobri o significado de underline, serve para ignorar um dos retornos
	fmt.Println(resultadoSoma2)                      // 25
}
