package json_unmarshall

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/douglasandradeee/GoExercicies/advanced/json_unmarshall/model"
)

func BlogPost() {
	cliente := &http.Client{
		Timeout: time.Second * 30,
	}

	request, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/posts/1", nil)
	if err != nil {
		fmt.Println("[main] Erro ao criar um request para o Servidor. Erro: ", err.Error())
		return
	}
	request.SetBasicAuth("user", "password")
	response, err := cliente.Do(request)
	if err != nil {
		fmt.Println("[main] Erro ao abrir a página do Servidor. Erro: ", err.Error())
		return
	}
	defer response.Body.Close()

	if response.StatusCode == 200 || response.StatusCode == 201 {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("[main] Erro ao ler o conteúdo da página do Servidor. Erro: ", err.Error())
			return
		}
		fmt.Println("")
		post := model.BlogPost{}
		err = json.Unmarshal(body, &post)
		if err != nil {
			fmt.Println("[main] Erro ao converter retorno JSON do Servidor. Erro: ", err.Error())
			return
		}
		fmt.Printf("Post é: %v\r\n", post)
	}
}
