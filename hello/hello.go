package main

import "fmt"
import "os"
import "net/http"
import "time"
import "strings"
import "bufio"


const monitoringTimes int = 3
const monitoringDelay = 5

func main() {
	showIntroduction()
	readSitesOfFile()

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
	version := 1.1

	fmt.Println("Olá, sr.", name)
	fmt.Println("Este programa está na versão", version)
}

func showMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func readCommand() int {
	var readComannd int
	fmt.Scan(&readComannd)
	fmt.Println("O comando escolhido foi", readComannd)
	return readComannd
}

func startMonitoring() {
	fmt.Println("Monitorando...")

	sites := readSitesOfFile()

	for i := 0; i < monitoringTimes; i++ {
		for i, site := range sites {
			fmt.Println("Testando site: ", i, " ", site)
			testSite(site)
		}
		time.Sleep(monitoringDelay * time.Second)
	}

	fmt.Println("")

}

func testSite(site string) {
	response, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	if response.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "esta com problemas. Status Code: ", response.StatusCode)
	}
}

func readSitesOfFile() []string {
    var sites []string
    file, err := os.Open("sites.txt")

    if err != nil {
        fmt.Println("Ocorreu um erro: ", err)
        return sites
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        site := strings.TrimSpace(scanner.Text())
        sites = append(sites, site)
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Ocorreu um erro: ", err)
    }

    return sites
}
