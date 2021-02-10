package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/douglasandradeee/GoExercicies/web_server/model"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	hour := time.Now().Format("15:04:05")
	page := model.Page{}
	page.Hour = hour
	if err := Model.ExecuteTemplate(w, "hello.html", page); err != nil {
		http.Error(w, "Houve um erro na renderização da página", http.StatusInternalServerError)
		fmt.Println("[Hello] Erro na execução do modelo. Erro: ", err.Error())
	}
}
