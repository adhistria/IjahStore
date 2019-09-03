package controller

import (
	"net/http"

	"context"
	"encoding/json"

	"github.com/adhistria/ijahstore/model"
	"github.com/adhistria/ijahstore/response"
	"github.com/adhistria/ijahstore/service"
	"github.com/gorilla/mux"
)

type productController struct {
	productService service.ProductService
}

func (pc *productController) AddProduct(w http.ResponseWriter, r *http.Request) {

	var product model.Product
	err := json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		response.APIErrorResponse(w, 500, err.Error())
		return
	}

	_, err = pc.productService.AddProduct(context.Background(), &product)

	if err != nil {
		response.APIErrorResponse(w, 500, err.Error())
		return
	}

	response.APISuccessResponse(w, 200, product)
	return
}

func (pc *productController) GetProducts(w http.ResponseWriter, r *http.Request) {

	products, err := pc.productService.GetProducts(context.Background())

	if err != nil {
		response.APIErrorResponse(w, 500, err.Error())
		return
	}

	response.APISuccessResponse(w, 200, products)
	return
}

func NewProductController(router *mux.Router, ps service.ProductService) productController {
	controller := productController{productService: ps}
	router.HandleFunc("/products", controller.AddProduct).Methods("POST")
	router.HandleFunc("/products", controller.GetProducts).Methods("GET")
	return controller
}
