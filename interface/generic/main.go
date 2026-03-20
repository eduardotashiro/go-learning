package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
	fmt.Println()
}

func main() {
	generica("String")
	generica(1)
	generica(true)

	fmt.Println(1, 2, "string", false, true, float64(1234))

	mapa := map[interface{}]interface{}{
		1:            "string",
		float32(123): true,
		"string":     "string",
		true:         13,
	}

	fmt.Println(mapa)
}


//QUE LOUCURA