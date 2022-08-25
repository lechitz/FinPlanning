package modelos

import "time"

type Item struct {
	ID                 uint64    `json:"id,omitempty"`
	Item               string    `json:"item"`
	Valor              uint64    `json:"valor"`
	QuantidadeDeParcelas uint64    `json:"quantidadeDeParcelas"`
	Beneficiario       string    `json:"beneficiario"`
	CompradoEm         time.Time `json:"compradoEm,omitempty"`
}
