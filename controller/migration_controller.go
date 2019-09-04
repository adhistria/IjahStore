package controller

import (
	"net/http"

	"github.com/adhistria/ijahstore/response"
	"github.com/adhistria/ijahstore/service"
	"github.com/gorilla/mux"
)

type migrationController struct {
	migrationService service.MigrationService
}

func (rc *migrationController) MigrateProducts(w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(32 << 20)

	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		response.APIErrorResponse(w, 500, err.Error())
		return
	}

	err = rc.migrationService.MigrateProducts(file, handler)
	if err != nil {
		response.APIErrorResponse(w, 500, err.Error())
		return
	}

	response.APISuccessResponse(w, 200, "Success Upload Data")
	return
}

func (rc *migrationController) MigrateIncomingProducts(w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(32 << 20)

	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		response.APIErrorResponse(w, 500, err.Error())
		return
	}

	err = rc.migrationService.MigrateIncomingProducts(file, handler)
	if err != nil {
		response.APIErrorResponse(w, 500, err.Error())
		return
	}

	response.APISuccessResponse(w, 200, "Success Upload Data")
	return
}

func (rc *migrationController) MigrateOutgoingProducts(w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(32 << 20)

	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		response.APIErrorResponse(w, 500, err.Error())
		return
	}

	err = rc.migrationService.MigrateOutgoingProducts(file, handler)

	if err != nil {
		response.APIErrorResponse(w, 500, err.Error())
		return
	}

	response.APISuccessResponse(w, 200, "Success Upload Data")
}

func NewMigrationController(router *mux.Router, ms service.MigrationService) migrationController {
	controller := migrationController{migrationService: ms}
	router.HandleFunc("/migrate-data-products", controller.MigrateProducts).Methods("POST")
	router.HandleFunc("/migrate-data-incoming-products", controller.MigrateIncomingProducts).Methods("POST")
	router.HandleFunc("/migrate-data-outgoing-products", controller.MigrateOutgoingProducts).Methods("POST")

	return controller
}
