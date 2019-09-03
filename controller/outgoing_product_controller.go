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

type outgoingProductController struct {
	OutgoingProductService service.OutgoingProductService
}

func (opc *outgoingProductController) AddOutgoingProduct(w http.ResponseWriter, r *http.Request) {

	var OutgoingProduct model.OutgoingProduct

	err := json.NewDecoder(r.Body).Decode(&OutgoingProduct)
	if err != nil {
		response.APIErrorResponse(w, 500, err.Error())
		return
	}

	_, err = opc.OutgoingProductService.AddOutgoingProduct(context.Background(), &OutgoingProduct)

	if err != nil {
		response.APIErrorResponse(w, 500, err.Error())
		return
	}

	response.APISuccessResponse(w, 200, OutgoingProduct)
	return
}

func NewOutgoingProductController(router *mux.Router, ips service.OutgoingProductService) outgoingProductController {
	controller := outgoingProductController{OutgoingProductService: ips}

	router.HandleFunc("/outgoing-products", controller.AddOutgoingProduct).Methods("POST")
	return controller
}
