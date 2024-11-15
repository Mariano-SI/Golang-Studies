package main

import "fmt"
import "os"
import "net/http"
import "time"

const monitoringTimes = 3
const monitoringDelay = 5

func main() {
	showIntroduction()

	for {
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
	
	sites := []string{"https://marianosilva.dev.br/", "https://www.alura.com.br", "https://www.caelum.com.br"}


	fmt.Println(sites)

	for i := 0; i < monitoringTimes; i++{
		for i, site := range sites{
			fmt.Println("Testando site: ", i, " ", site)
			testSite(site)
		}
		time.Sleep(monitoringTimes * time.Second)
	}


	fmt.Println("")

}

func testSite(site string){
	response, _ := http.Get(site)
	if response.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	}else{
		fmt.Println("Site:", site, "esta com problemas. Status Code: ", response.StatusCode)
	}	
}

