package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)
import (
	"../repository"
	"../utils"
	"../models"
	)

type ProductsController interface {
	PostProduct(http.ResponseWriter, *http.Request)
}

type productsControllerImpl struct {
	productsRepository repository.ProductsRepository
}

func NewProductsController(productsRepository repository.ProductsRepository) *productsControllerImpl {
	return &productsControllerImpl{productsRepository}
}

func (c *productsControllerImpl) PostProduct(w http.ResponseWriter, r *http.Request) {
	if r.Body != nil {
		defer r.Body.Close()
	}

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	product := &models.Product{}
	err = json.Unmarshal(bytes, product)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	err = product.Validate()
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	product, err = c.productsRepository.Save(product)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	buildCreatedResponse(w, buildLocation(r, product.ID))
	utils.WriteAsJson(w, product)
}