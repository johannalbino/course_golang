package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("String")
	generica(15)
	generica(true)

	mapa := map[interface{}]interface{}{
		1:            "String",
		float64(100): true,
		"String":     "String",
		true:         float64(12),
	}

	fmt.Println(mapa)
}
