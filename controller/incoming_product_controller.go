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

type incomingProductController struct {
	incomingProductService service.IncomingProductService
}

func (ipc *incomingProductController) AddIncomingProduct(w http.ResponseWriter, r *http.Request) {

	var incomingProduct model.IncomingProduct

	err := json.NewDecoder(r.Body).Decode(&incomingProduct)
	if err != nil {
		response.APIErrorResponse(w, 500, err.Error())
		return
	}

	_, err = ipc.incomingProductService.AddIncomingProduct(context.Background(), &incomingProduct)

	if err != nil {
		response.APIErrorResponse(w, 500, err.Error())
		return
	}

	response.APISuccessResponse(w, 200, incomingProduct)
	return
}

func (ipc *incomingProductController) GetIncomingProducts(w http.ResponseWriter, r *http.Request) {

	incomingProducts, err := ipc.incomingProductService.GetIncomingProducts(context.Background())

	if err != nil {
		response.APIErrorResponse(w, 500, err.Error())
		return
	}

	response.APISuccessResponse(w, 200, incomingProducts)
	return
}

func NewIncomingProductController(router *mux.Router, ips service.IncomingProductService) incomingProductController {
	controller := incomingProductController{incomingProductService: ips}

	router.HandleFunc("/incoming-products", controller.AddIncomingProduct).Methods("POST")
	router.HandleFunc("/incoming-products", controller.GetIncomingProducts).Methods("GET")
	return controller
}
