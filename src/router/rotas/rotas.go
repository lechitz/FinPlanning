package rotas

import (
	"github.com/gorilla/mux"
	"net/http"
)

// Rota representa todas as rotas da API
type Rota struct {
	URI    string
	Metodo string
	Funcao func(http.ResponseWriter, *http.Request)
}

//Configurar Retorna o "router" com todas as rotas já configuradas
func Configurar(r *mux.Router) *mux.Router {

	rotas := rotasItens
	rotas = append(rotas, rotaPrincipal)

	for _, rota := range rotas {
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}
	return r
}
