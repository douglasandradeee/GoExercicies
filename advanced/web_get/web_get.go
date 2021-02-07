package web_get

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func Get() {
	cliente := &http.Client{
		Timeout: time.Second * 30,
	}
	response, err := cliente.Get("https://www.google.com.br")
	if err != nil {
		fmt.Println("[main] Erro ao abrir a página do google Brasil. Erro: ", err.Error())
		return
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("[main] Erro ao ler o conteúdo da página do google Brasil. Erro: ", err.Error())
			return
		}
		fmt.Println(string(body))
	}

	request, err := http.NewRequest("GET", "https://www.google.com.br", nil)
	if err != nil {
		fmt.Println("[main] Erro ao criar um request para o google Brasil. Erro: ", err.Error())
		return
	}
	request.SetBasicAuth("user", "password")
	response, err = cliente.Do(request)
	if err != nil {
		fmt.Println("[main] Erro ao abrir a página do google Brasil. Erro: ", err.Error())
		return
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("[main] Erro ao ler o conteúdo da página do google Brasil. Erro: ", err.Error())
			return
		}
		fmt.Println("")
		fmt.Println(string(body))
	}
}
