package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Israel-Ferreira/shopping-api/internal/services"
	"github.com/Israel-Ferreira/shopping-api/pkg/models"
)

type ShoppingController struct {
	Service services.ShoppingService
}

func (sc ShoppingController) ListShoppings(w http.ResponseWriter, r *http.Request) {
	shoppings, err := sc.Service.FindAll()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(&shoppings); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func (sc ShoppingController) GetShoppingById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	shopping, err := sc.Service.FindById(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(&shopping); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func (sc ShoppingController) AddShopping(w http.ResponseWriter, r *http.Request) {
	var shopping models.ShoppingRequest

	if err := json.NewDecoder(r.Body).Decode(&shopping); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := sc.Service.CreateShopping(shopping); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	msg := map[string]string{
		"msg": "",
	}

	if err := json.NewEncoder(w).Encode(&msg); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func (sc ShoppingController) DeleteShopping(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if err := sc.Service.DeleteById(id); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (sc ShoppingController) UpdateShopping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}
