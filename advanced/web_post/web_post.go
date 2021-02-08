package web_post

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/douglasandradeee/GoExercicies/advanced/web_post/model"
)

func Post() {
	cliente := &http.Client{
		Timeout: time.Second * 30,
	}

	user := model.User{}
	user.Id = 1
	user.Name = "Gabriela do Nada"

	sendContent, err := json.Marshal(&user)
	if err != nil {
		fmt.Println("[main] Erro ao gerar o JSON do objeto usuário. Erro: ", err.Error())
		return
	}

	request, err := http.NewRequest("POST", "https://requestbin.net/r/ecz2e5hw", bytes.NewBuffer(sendContent))
	if err != nil {
		fmt.Println("[main] Erro ao criar um request para o requestBin. Erro: ", err.Error())
		return
	}
	request.SetBasicAuth("fizz", "buzz")
	request.Header.Set("content-type", "application/json; charset=utf-8")
	response, err := cliente.Do(request)
	if err != nil {
		fmt.Println("[main] Erro ao chamar o serviço do requestBin. Erro: ", err.Error())
		return
	}
	defer response.Body.Close()

	if response.StatusCode == 200 || response.StatusCode == 201 {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("[main] Erro ao ler o conteúdo pela requestBin. Erro: ", err.Error())
			return
		}
		fmt.Println(" ")
		fmt.Println(string(body))
	}
}
