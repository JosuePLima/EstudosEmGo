package main

import "fmt"

func main() {
	var variavel1 string = "Variável"
	variavel2 := "Variável 2"

	variavel1, variavel2 = variavel2, variavel1

	fmt.Println(variavel1)
	fmt.Println(variavel2)

}