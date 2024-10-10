package services

import (
	"github.com/Israel-Ferreira/shopping-api/pkg/interfaces"
	"github.com/Israel-Ferreira/shopping-api/pkg/models"
)

type ShoppingService struct {
	DbRepo interfaces.DbRepository
}

func (shoppingService *ShoppingService) FindAll() ([]models.Shopping, error) {
	shoppings, err := shoppingService.DbRepo.ListShoppings()

	if err != nil {
		return []models.Shopping{}, err
	}

	return shoppings, nil
}

func (shoppingService *ShoppingService) FindById(id string) (models.Shopping, error) {
	shopping, err := shoppingService.DbRepo.GetShoppingById(id)

	if err != nil {
		return models.Shopping{}, err
	}

	return shopping, nil
}

func (shoppingService *ShoppingService) DeleteById(id string) error {

	if err := shoppingService.DbRepo.DeleteShoppingById(id); err != nil {
		return err
	}

	return nil
}

func (shoppingService *ShoppingService) CreateShopping(dto models.ShoppingRequest) error {
	if err := shoppingService.DbRepo.CreateShopping(dto); err != nil {
		return err
	}

	return nil
}
