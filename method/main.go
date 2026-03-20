package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() { // 'u' pode ser qualquer variavel, geralmente é a primeira letra da struct
	fmt.Printf("Salvando os dados do Usuário %s no banco de dados\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Usuário 1", 17}
	fmt.Println(usuario1) //{Usuário 1 17}
	usuario1.salvar()     //Salvando os dados do Usuário Usuário 1 no banco de dados

	usuario2 := usuario{"Davi", 30}
	usuario2.salvar() // Salvando os dados do Usuário Davi no banco de dados

	maiorDeIdade := usuario2.maiorDeIdade()
	fmt.Println(maiorDeIdade) //true

	fmt.Println(usuario2.idade) // 30
	fmt.Println(usuario2.idade) // 30
	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade) // 31
	fmt.Println(usuario2.idade) // 31
}
