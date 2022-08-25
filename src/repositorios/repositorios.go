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

func (repositorio Itens) BuscarItens(qualItem string) ([]modelos.Item, error) {
	qualItem = fmt.Sprintf("%%%s%%", qualItem)

	linhas, erro := repositorio.db.Query("SELECT id, item, valor, quantidadeDeParcelas, beneficiario, compradoEm from itens where item LIKE ?", qualItem)
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