package main

import (
	"errors"
	"fmt"
)

// func main()  {
// 	// int8, int16, int32, int64
// 	var numero int16 = 100
// 	fmt.Println(numero)
// }

// func main() {
// 	// int8, int16, int32, int64
// 	var numero int16 = 10000000 //cannot use 10000000 (untyped int constant) as int16 value in variable declaration (overflows)compilerNumericOverflow
// 	fmt.Println(numero)
// }

// func main()  {
// 	// int8, int16, int32, int64
// 	var numero int = 100 // usa a arquitetura do meu PC como base
// 	fmt.Println(numero)
// }

// func main()  {
// 	// int8, int16, int32, int64
// 	 numero := 100  // var numero int //usa a arquitetura do meu PC como base
// 	fmt.Println(numero)
// }

func main() {
	// int8, int16, int32, int64
	var numero int16 = 100
	fmt.Println(numero)

	// uint8, uint16, uint32, uint64
	var num2 uint64 = 111
	fmt.Println(num2)

	//*/* rune is an alias for int32 and is equivalent to int32 in all ways. It is
	//*/* used, by convention, to distinguish character values from integer values.
	//*/* type rune = int32
	var num3 rune = 222
	fmt.Println(num3)

	//*/* byte is an alias for uint8 and is equivalent to uint8 in all ways. It is
	//*/* used, by convention, to distinguish byte values from 8-bit unsignedinteger values.
	//*/*type byte = uint8
	var num4 byte = 222
	fmt.Println(num4)

	/*/*/ /*/*/ /*/*/ /*/*/ /*//*/ //

	//FLOAT

	var numeroReal1 float32 = 123.456
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123.456
	fmt.Println(numeroReal2)

	numeroReal3 := 123.456
	fmt.Println(numeroReal3)

	//STRING

	var str string = "texto"
	fmt.Println(str)

	str2 := "texto2"
	fmt.Println(str2)

	char := 't' 
	fmt.Println(char) // output -> 116, numero que a letra t esta na tabela acii

	//VALOR ZERO

	var str3 string
	fmt.Println(str3) // 

	//BOOLEAN

	var tv bool = false
	fmt.Println(tv) // como alguém diz que null é igual a false...kkk esse mico vou guardar pra vida kk

	tv1 := true
	fmt.Println(tv1)

	var tv2 bool
	fmt.Println(tv2) //false

	//ERROR   //tema bateu muito com o clima acima

	var erro1 error 
	fmt.Println(erro1) //<nil>

	var erro2 error = errors.New("TRATANDO ERROS")
	fmt.Println(erro2) 

}

// uint em Go é um tipo de dado inteiro sem sinal (unsigned integer),
// o que significa que ele armazena apenas números positivos e zero,
// dobrando a capacidade de armazenamento positivo em comparação ao int.
//  tamanho do uint depende da arquitetura (32 bits em sistemas 32-bit, 64 bits em 64-bit) , não pesquisei no google, jamais...
