package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func AdicionarItem(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Adicionando um novo item"))
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil{
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var item modelos.Item
	if erro = json.Unmarshal(corpoRequest, &item); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	//Criando o repositório
	repositorio := repositorios.NovoRepositorioDeItens(db)
	//Chamando o usuário que tá vindo na requisição
	item.ID, erro = repositorio.AdicionarItem(item)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, item)
}

func BuscarItens(w http.ResponseWriter, r *http.Request) {
	buscandoItens := strings.ToLower(r.URL.Query().Get("item"))

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeItens(db)
	itens, erro := repositorio.BuscarItens(buscandoItens)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, itens)
}

func BuscarItem(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	itemID, erro := strconv.ParseUint(parametros["itemId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeItens(db)
	item, erro := repositorio.BuscarItemID(itemID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, item)
}

func AtualizarItem(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	itemID, erro := strconv.ParseUint(parametros["itemId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var item modelos.Item
	if erro = json.Unmarshal(corpoRequisicao, &item); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeItens(db)
	if erro = repositorio.AtualizarItem(itemID, item); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

func DeletarItem(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando um item"))
}