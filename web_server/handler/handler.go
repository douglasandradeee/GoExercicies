package handler

import (
	"fmt"
	"net/http"
)

func Notice(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Aqui é um manipulador usando função num pacote")
}
