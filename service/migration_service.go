package service

import (
	"encoding/csv"
	"fmt"
	"log"
	"mime/multipart"
	"strconv"
	"strings"
	"time"

	"github.com/adhistria/ijahstore/model"
	"github.com/adhistria/ijahstore/repository"
)

type (
	MigrationService interface {
		MigrateProducts(multipart.File, *multipart.FileHeader) error
		MigrateIncomingProducts(multipart.File, *multipart.FileHeader) error
	}

	migrationServiceImpl struct {
		productRepository         repository.ProductRepository
		incomingProductRepository repository.IncomingProductRepository
	}
)

func (ms *migrationServiceImpl) MigrateProducts(file multipart.File, handler *multipart.FileHeader) error {

	var products []model.Product
	defer file.Close()

	fmt.Println("disini")
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Print(err)
		return err
	}

	for _, record := range records[1:] {
		total, err := strconv.Atoi(record[2])
		if err != nil {
			log.Print(err)
			return err
		}
		product := model.Product{
			SKU:   record[0],
			Name:  record[1],
			Total: total,
		}
		products = append(products, product)
		err = ms.productRepository.Add(&product)
	}

	fmt.Println("diakhir")
	fmt.Println(err)
	return err

}

func (ms *migrationServiceImpl) MigrateIncomingProducts(file multipart.File, handler *multipart.FileHeader) error {

	var incomingProducts []model.IncomingProduct
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Print(err)
		return err
	}

	for _, record := range records[1:] {

		fixDate := record[0] + ":00.000+07:00"

		date, err := time.Parse("2006/1/02 15:04:05.000-07:00", fixDate)

		if err != nil {
			return err
		}

		totalOrder, err := strconv.Atoi(record[3])
		if err != nil {
			return err
		}

		totalReceiveOrder, err := strconv.Atoi(record[4])
		if err != nil {
			return err
		}

		purchasePriceString := strings.Replace(record[5][2:], ".", "", -1)
		purchasePrice, err := strconv.Atoi(purchasePriceString)
		if err != nil {
			return err
		}

		totalPurchasePriceString := strings.Replace(record[6][2:], ".", "", -1)
		totalPurchasePrice, err := strconv.Atoi(totalPurchasePriceString)
		if err != nil {
			return err
		}
		incomingProduct := model.IncomingProduct{
			CreatedAt:          date,
			ProductID:          record[1],
			TotalOrder:         totalOrder,
			TotalReceiveOrder:  totalReceiveOrder,
			PurchasePrice:      purchasePrice,
			TotalPurchasePrice: totalPurchasePrice,
			ReceiptNumber:      record[7],
			Notes:              record[8],
		}

		incomingProducts = append(incomingProducts, incomingProduct)
		err = ms.incomingProductRepository.AddWithTimestamps(&incomingProduct)
		if err != nil {
			log.Print(err)
			return err
		}
		err = ms.productRepository.AddTotalProduct(&incomingProduct)
	}
	return err
}

func NewMigrationService(pr repository.ProductRepository, ipr repository.IncomingProductRepository) MigrationService {
	return &migrationServiceImpl{
		productRepository:         pr,
		incomingProductRepository: ipr,
	}
}
