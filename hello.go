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

const monitorings = 3
const delay = 5

func main() {
	for {
		fmt.Println("1 - Iniciar monitoramento")
		fmt.Println("2 - Exibir Logs")
		fmt.Println("0 - Sair do Programa")

		switch handleCommand() {
		case 1:
			startMonitoring()
		case 2:
			showLogs()
		case 0:
			fmt.Println("Saindo do programa")
		default:
			fmt.Println("Não conheço este comando")
		}
	}

}

func handleCommand() int {
	var command int
	fmt.Scan(&command)

	return command
}

func startMonitoring() {
	fmt.Println("Monitorando")

	sites := readSitesFile()
	for i := 0; i < monitorings; i++ {
		for _, site := range sites {
			testUrl(site)
		}
		time.Sleep(delay * time.Second)

	}

}

func testUrl(url string) {
	response, err := http.Get(url)

	if err != nil {
		fmt.Println("Houve um erro:", err)
	}
	if response.StatusCode == 200 {
		fmt.Println("Site:", url, "foi carregado com sucesso!")
		saveLog(url, true)
	} else {
		fmt.Println("Site: ", url, "esta com problemas, status code:", response.StatusCode)
		saveLog(url, false)
	}
}

func readSitesFile() []string {
	var sites []string
	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		sites = append(sites, line)

		if err == io.EOF {
			break
		}

	}

	file.Close()
	return sites
}

func saveLog(site string, status bool) {

	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "- online: " + strconv.FormatBool(status) + "\n")
	file.Close()
}

func showLogs() {
	file, err := os.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(file))
}
