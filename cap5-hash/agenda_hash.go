package main

import "fmt"

func main() {

	agenda := map[string]int{
		"Tata": 9999,
		"Anna": 8888,
		"Tete": 7777,
	}

	fmt.Println("Telefone da Anna:", agenda["Anna"])
}
