package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/douglasandradeee/GoExercicies/writeFiles/model"
)

func main() {
	file, err := os.Open("city.csv")
	if err != nil {
		fmt.Println("[main] Houve um erro ao abrir o arquivo. Erro:", err.Error())
		return
	}
	defer file.Close()

	readerCsv := csv.NewReader(file)
	content, err := readerCsv.ReadAll()
	if err != nil {
		fmt.Println("[main] Houve um erro ao ler o arquivo com leitor CSV. Erro:", err.Error())
		return
	}

	fileJSON, err := os.Create("city.json")
	if err != nil {
		fmt.Println("[main] Houve um erro ao criar o arquivo JSON. Erro:", err.Error())
		return
	}
	defer fileJSON.Close()

	writer := bufio.NewWriter(fileJSON)
	writer.WriteString("[\r\n")

	for _, line := range content {
		for indexItem, item := range line {
			data := strings.Split(item, "/")
			city := model.City{}
			city.Name = data[0]
			city.State = data[1]
			fmt.Printf("City: %+v\r\n", city)
			cityJSON, err := json.Marshal(city)
			if err != nil {
				fmt.Println("[main] Houve um erro ao gerar o JSON do Item ", item, ". Erro: ", err.Error())
			}
			writer.WriteString(" " + string(cityJSON))
			if (indexItem + 1) < len(line) {
				writer.WriteString(",\r\n")
			}
		}
	}
	writer.WriteString("\r\n]")
	writer.Flush()
}
