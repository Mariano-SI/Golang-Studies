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


	if command == 1 {
		fmt.Println("Monitorando...")
	} else if command == 2 {
		fmt.Println("Exibindo Logs...")
	} else if command == 0 {
		fmt.Println("Saindo do Programa...")
	} else {
		fmt.Println("Não conheço este comando")
	}
	
	switch command {
		case 1:
			fmt.Println("Monitorando...")
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do Programa...")
		default:
			fmt.Println("Não conheço este comando")
	}
	
}