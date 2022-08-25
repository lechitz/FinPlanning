package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotaPrincipal = Rota{
	URI: "/itens",
	Metodo: http.MethodPost,
	Funcao: controllers.Principal,
}
