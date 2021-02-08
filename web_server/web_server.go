package web_server

import (
	"fmt"
	"net/http"

	"github.com/douglasandradeee/GoExercicies/web_server/handler"
)

func ServerTest() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hellow World!!")
	})
	http.HandleFunc("/handleNotice", handler.Notice)
	fmt.Println("Servidor iniciado ....")
	http.ListenAndServe(":8181", nil)
}
