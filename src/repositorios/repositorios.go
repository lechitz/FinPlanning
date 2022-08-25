package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

//Itens representa um repositório de itens
type Itens struct {
	db *sql.DB
}

func NovoRepositorioDeItens(db *sql.DB) *Itens {
	return &Itens{db}
}

func (repositorio Itens) AdicionarItem(item modelos.Item) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into itens (item, valor, quantidadeDeParcelas, beneficiario) values(?,?,?,?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(item.Item, item.Valor, item.QuantidadeDeParcelas, item.Beneficiario)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

func (repositorio Itens) BuscarItens(todosItens string) ([]modelos.Item, error) {
	todosItens = fmt.Sprintf("%%%s%%", todosItens)

	linhas, erro := repositorio.db.Query("SELECT id, item, valor, quantidadeDeParcelas, beneficiario, compradoEm from itens where item LIKE ?", todosItens)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var itens []modelos.Item

	for linhas.Next() {
		var item modelos.Item

		if erro = linhas.Scan(
			&item.ID,
			&item.Item,
			&item.Valor,
			&item.QuantidadeDeParcelas,
			&item.Beneficiario,
			&item.CompradoEm,
			); erro != nil {
			return nil, erro
			}

			itens = append(itens, item)
	}

	return itens, nil
}

func (repositorio Itens) BuscarItemID(qualItem uint64) (modelos.Item, error) {
	linhas, erro := repositorio.db.Query("select id, item, valor, quantidadeDeParcelas, beneficiario, compradoEm from itens where id = ?", qualItem)
	if erro != nil {
		return modelos.Item{}, erro
	}
	defer linhas.Close()

	var item modelos.Item

	if linhas.Next() {
		if erro = linhas.Scan(
			&item.ID,
			&item.Item,
			&item.Valor,
			&item.QuantidadeDeParcelas,
			&item.Beneficiario,
			&item.CompradoEm,
			); erro != nil {
			return modelos.Item{}, erro
		}
	}

	return item, nil
}

func (repositorio Itens) AtualizarItem(ID uint64, item modelos.Item) error {
	statement, erro := repositorio.db.Prepare("update itens set item=?, valor=?, quantidadeDeParcelas=?, beneficiario=? where id=?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(item.Item, item.Valor, item.QuantidadeDeParcelas, item.Beneficiario, ID); erro != nil {
		return erro
	}
	return nil
}

func (repositorio Itens) DeletarItem(ID uint64) error {
	statement, erro := repositorio.db.Prepare("delete from itens where id=?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}