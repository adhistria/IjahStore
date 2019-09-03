package main

import (
	"log"
	"net/http"

	"github.com/adhistria/ijahstore/controller"
	"github.com/adhistria/ijahstore/driver"
	"github.com/adhistria/ijahstore/repository"
	"github.com/adhistria/ijahstore/service"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	db := driver.GetConnection()
	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	incomingProductRepository := repository.NewIncomingProductRepository(db)
	incomingProductService := service.NewIncomingProductService(productRepository, incomingProductRepository)
	outgoingProductRepository := repository.NewOutgoingProductRepository(db)
	outgoingProductService := service.NewOutgoingProductService(productRepository, outgoingProductRepository)
	controller.NewProductController(router, productService)
	controller.NewIncomingProductController(router, incomingProductService)
	controller.NewOutgoingProductController(router, outgoingProductService)
	log.Print("Starting Server")
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatal("Port 8000 Failed, connection refused")
	}

}
