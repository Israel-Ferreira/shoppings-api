package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Endereco struct {
	Logradouro string  `json:"logradouro"`
	Bairro     string  `json:"bairro"`
	Cidade     string  `json:"cidade"`
	Estado     string  `json:"estado"`
	Telefone   string  `json:"telefone"`
	Numero     string  `json:"numero"`
	Lat        float64 `json:"latitude"`
	Long       float64 `json:"longitude"`
}

type Shopping struct {
	Id                  primitive.ObjectID `json:"_id"`
	Nome                string             `json:"nome"`
	QtdeSalasCinema     uint               `json:"qtde_salas_cinema"`
	QtdeLojas           uint               `json:"qtde_lojas"`
	Tipo                string             `json:"tipo"`
	TemPracaAlimentacao bool               `json:"tem_praca_alimentacao"`
	Endereco            Endereco           `json:"endereco"`
}

func NewShopping(shoppingReq ShoppingRequest) *Shopping {
	return &Shopping{
		Nome:                shoppingReq.Nome,
		Tipo:                shoppingReq.Tipo,
		QtdeSalasCinema:     shoppingReq.QtdeSalasCinema,
		TemPracaAlimentacao: shoppingReq.TemPracaAlimentacao,
		QtdeLojas:           shoppingReq.QtdeLojas,
		Endereco: Endereco{
			Logradouro: shoppingReq.Endereco.Logradouro,
			Bairro:     shoppingReq.Endereco.Bairro,
			Cidade:     shoppingReq.Endereco.Cidade,
			Estado:     shoppingReq.Endereco.Estado,
			Lat:        shoppingReq.Endereco.Lat,
			Long:       shoppingReq.Endereco.Long,
			Numero:     shoppingReq.Endereco.Numero,
		},
	}
}
