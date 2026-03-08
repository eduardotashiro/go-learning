package main //o que importa é o nome do pacote, pode ser diferente do arquivo

import (
	"fmt"
)

func main() {
	var variavel1 string = "variavel1" //EXPLÍCITO
	variavel2 := "variável2"           //variables.go:9:2: declared and not used: variavel2 #NÃO PODE DECLARAR E NÃO USAR
	fmt.Println(variavel1)
	// fmt.Println(variavel2)
	fmt.Println(variavel2)

	var (
		variavel3 string = "Ken Thompson"
		variavel4 string = "\nEduardo Tashiro"
	)
	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "variavel5", "\nvariavel6"
	fmt.Println(variavel5, variavel6)

	const constante1 string = "constante1"
	fmt.Println(constante1)

	const (
		constante2 string = "constante2\n"
		constante3 string = "constante3"
	)
	fmt.Print(constante2, constante3)

	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)

}
