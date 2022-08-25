package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasItens = []Rota{
	{
		//Método POST para produto
		URI:    "/itens",
		Metodo: http.MethodPost,
		Funcao: controllers.AdicionarItem,
	},
	{
		//Método GET para todos os itens
		URI: "/itens",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarItens,
	},
	{
		//Método GET para um item em específico
		URI: "/itens/{itemId}",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarItem,
	},
	{
		//Método PUT para um item em específico
		URI: "/itens/{itemId}",
		Metodo: http.MethodPut,
		Funcao: controllers.AtualizarItem,
	},
	{
		//Método DELETE para um item específico
		URI: "/itens/{itemId}",
		Metodo: http.MethodDelete,
		Funcao: controllers.DeletarItem,
	},
}
