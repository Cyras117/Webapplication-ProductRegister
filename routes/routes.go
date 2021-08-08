package routes

import (
	c "go_cadastro/controllers"
	"net/http"
)

func CarregaRotas() {
	http.HandleFunc("/", c.Index)
	http.HandleFunc("/new", c.New)
}
