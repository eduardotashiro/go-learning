package main

import "fmt"

func funcao1() {
	fmt.Println("executando a funcao 1")
}

func funcao2() {
	fmt.Println("executando a funcao 2")
}
// usar no PGSQL para fechar conexão com o banco..
func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada, resultado será retornado")
	fmt.Println("entrando na função para verificar se o aluno esta aprovado")
	media := (n1 + n2) / 2
	if media >= 6 {
		return true
	}

	return false
}

func main() {
	defer funcao1()
	// DEFER = ADIAR
	funcao2()

	fmt.Println(alunoEstaAprovado(1, 1))

}
