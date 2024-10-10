package routes

import (
	"net/http"

	"github.com/Israel-Ferreira/shopping-api/internal/controllers"
	"github.com/Israel-Ferreira/shopping-api/internal/databases"
	"github.com/Israel-Ferreira/shopping-api/internal/services"
)

func InitRoutes() {
	mongoRepo := databases.NewRepository(databases.Client)
	service := services.ShoppingService{
		DbRepo: mongoRepo,
	}

	controller := controllers.ShoppingController{
		Service: service,
	}

	http.HandleFunc("GET /shoppings", controller.ListShoppings)
	http.HandleFunc("POST /shoppings", controller.AddShopping)

	http.HandleFunc("PUT /shoppings/{id}", controller.UpdateShopping)
	http.HandleFunc("DELETE /shoppings/{id}", controller.DeleteShopping)
	http.HandleFunc("GET /shoppings/{id}", controller.GetShoppingById)
}
