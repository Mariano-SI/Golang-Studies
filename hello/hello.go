package main

import "fmt"
import "reflect"

func main() {
	name := "Mariano"
	version :=  1.1
	age := 23
	fmt.Println("Olá, sr.", name, " sua idade é ", age)
	fmt.Println("Este programa está na versão", version)

	fmt.Println("O tipo da variável name é", reflect.TypeOf(version))
}