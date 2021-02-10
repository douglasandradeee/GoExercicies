package handler

import "html/template"

var Model = template.Must(template.ParseFiles("web_server/html/hello.html"))
