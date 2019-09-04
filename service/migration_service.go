package service

import (
	"encoding/csv"
	"fmt"
	"log"
	"mime/multipart"
	"strconv"

	"github.com/adhistria/ijahstore/model"
	"github.com/adhistria/ijahstore/repository"
)

type (
	MigrationService interface {
		MigrateProducts(multipart.File, *multipart.FileHeader) error
	}

	migrationServiceImpl struct {
		productRepository repository.ProductRepository
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

func NewMigrationService(pr repository.ProductRepository) MigrationService {
	return &migrationServiceImpl{
		productRepository: pr,
	}
}
