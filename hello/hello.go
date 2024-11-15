package main

import "fmt"
import "os"
import "net/http"


func main() {

	showIntroduction()

	showMenu()

	command := readCommand()
	
	switch command {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do Programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
	}
	
}

func showIntroduction() {
	name := "Mariano"
	version :=  1.1

	fmt.Println("Olá, sr.", name)
	fmt.Println("Este programa está na versão", version)
}

func showMenu(){
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func readCommand()int{
	var readComannd int
	fmt.Scan(&readComannd)
	fmt.Println("O comando escolhido foi", readComannd)
	return readComannd
}

func startMonitoring(){
	fmt.Println("Monitorando...")
	site := "https://marianosilva.dev.br/"
	response, _ := http.Get(site)

	if response.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	}
}
