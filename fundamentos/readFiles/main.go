package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("city.csv")
	if err != nil {
		fmt.Println("[main] Houve um erro ao abrir o arquivo. Erro:", err.Error())
		return
	}

	// scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// 	line := scanner.Text()
	// 	fmt.Println("O CONTEÚDO DA LINHA É: ", line)
	// }
	readerCsv := csv.NewReader(file)
	content, err := readerCsv.ReadAll()
	if err != nil {
		fmt.Println("[main] Houve um erro ao ler o arquivo com leitor CSV. Erro:", err.Error())
		return
	}
	for indexLine, line := range content {
		fmt.Printf("Line[%d] is %s\r\n", indexLine, line)
		for indexItem, item := range line {
			fmt.Printf("Item[%d] is %s\r\n", indexItem, item)
		}
	}

	file.Close()
}
