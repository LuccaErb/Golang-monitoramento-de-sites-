package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoring = 3
const delay = 5

func main() {

	displayIntroduction()

	for {
		dislpayMenu()

		command := readCommand()
		switch command {
		case 1:
			startingMonitoring()
		case 2:
			fmt.Println("Log")
			printLogs()
		case 0:
			fmt.Println("Saindo do Programa")
			os.Exit(0)
		default:
			fmt.Println("O comando não foi reconhecido")
			os.Exit(-1)

		}
	}

}
func displayIntroduction() {
	name := "Lucca"
	version := 1.1

	fmt.Println("hello, sir.", name)
	fmt.Println("The program in version: ", version)
}

func dislpayMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")

}
func readCommand() int {
	var commandRead int
	fmt.Scan(&commandRead)
	fmt.Println("O comando escolhido foi:", commandRead)

	return commandRead
}

func startingMonitoring() {
	fmt.Println("Monitorando")
	// sites := []string{"https://www.alura.com.br", "https://www.youtube.com", "https://www.youtube.com/watch?v=S1S1gp7UZNI"}

	sites := readSitesForFile()

	for n := 0; n < monitoring; n++ {
		for i, site := range sites {
			fmt.Println(i, site)
			testSite(site)
		}
		time.Sleep(delay * time.Second)
	}

}
func testSite(site string) {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}
	if resp.StatusCode == 200 {
		fmt.Println("site:", site, "foi carregado com sucesso, codigo de resposta:  ", resp.StatusCode)
		registerLog(site, true)
	} else {
		fmt.Println("site", site, "esta com problema status code : ", resp.StatusCode)
		registerLog(site, false)
	}
}
func readSitesForFile() []string {
	var sites []string
	file, err := os.Open("site.txt")
	// file, err := ioutil.ReadFile("site.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro :", err)
	}

	read := bufio.NewReader(file)
	for {
		line, err := read.ReadString('\n')
		line = strings.TrimSpace(line)
		sites = append(sites, line)
		if err == io.EOF {

			break

		}
	}

	file.Close()
	// fmt.Println(string(file))
	return sites
}
func registerLog(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	file.WriteString(time.Now().Format("02/01/2006 15:04:05 ") + site + " - online: " + strconv.FormatBool(status) + "\n")
	file.Close()
}
func printLogs() {
	file, err := os.ReadFile("log.txt")
	fmt.Println(string(file))
	if err != nil {
		fmt.Println("Nao foi possivel abrir o arquivo Log:", err)
	}
}
