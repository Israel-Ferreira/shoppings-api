package interfaces

import "github.com/Israel-Ferreira/shopping-api/pkg/models"

type DbRepository interface {
	ListShoppings() ([]models.Shopping, error)
	GetShoppingById(id string) (models.Shopping, error)
	DeleteShoppingById(id string) error
	UpdateShopping(id string, dtoReq models.ShoppingRequest) error
	CreateShopping(dtoReq models.ShoppingRequest) error
}
