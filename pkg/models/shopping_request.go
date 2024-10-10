package models

type ShoppingRequest struct {
	Nome                string   `json:"nome"`
	QtdeSalasCinema     uint     `json:"qtde_salas_cinema"`
	QtdeLojas           uint     `json:"qtde_lojas"`
	Tipo                string   `json:"tipo"`
	TemPracaAlimentacao bool     `json:"tem_praca_alimentacao"`
	Endereco            Endereco `json:"endereco"`
}
