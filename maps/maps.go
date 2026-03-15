package main

import (
	"fmt"
)

func main() {

	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}
	fmt.Println(usuario["sobrenome"]) //Silva

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "joao",
			"ultimo":   "pedro",
		},
		"curso": {
			"nome":   "ADS",
			"campus": "bloco 1",
		},
	}

	fmt.Println(usuario2) // map[curso:map[campus:bloco 1 nome:ADS] nome:map[primeiro:joao ultimo:pedro]] wtf
	delete(usuario2, "nome")
	fmt.Println(usuario2) // map[curso:map[campus:bloco 1 nome:ADS]]

	usuario2["signo"] = map[string]string{
		"nome": "virgem",
	}
	fmt.Println(usuario2) //map[curso:map[campus:bloco 1 nome:ADS] signo:map[nome:virgem]]

}
