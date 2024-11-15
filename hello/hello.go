package main

import "fmt"


func main() {
	name := "Mariano"
	version :=  1.1

	fmt.Println("Olá, sr.", name)
	fmt.Println("Este programa está na versão", version)

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")

	var command int
	fmt.Scan(&command)

	fmt.Println("O comando escolhido foi", command)
}