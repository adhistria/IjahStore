package repository

import (
	"github.com/adhistria/ijahstore/model"
	"github.com/jmoiron/sqlx"
)

type OutgoingProductRepository interface {
	FindAll() ([]model.OutgoingProduct, error)
	Add(*model.OutgoingProduct) error
	AddWithTimestamps(*model.OutgoingProduct) error
	GetSoldProducts(model.Date) ([]model.SoldProduct, error)
	GetTotalSoldProductByDate(model.Date) (int, error)
}

type outgoingProductRepositoryImpl struct {
	db *sqlx.DB
}

func (opr *outgoingProductRepositoryImpl) FindAll() ([]model.OutgoingProduct, error) {
	var outgoingProducts []model.OutgoingProduct
	err := opr.db.Select(&outgoingProducts, "SELECT * FROM outgoingProducts")
	return outgoingProducts, err
}

func (opr *outgoingProductRepositoryImpl) Add(op *model.OutgoingProduct) error {
	insertoutgoingProductQuery := `INSERT INTO OutgoingProducts (sold_amount, selling_price, total_selling_price, notes, product_id) VALUES( ?, ?, ?, ?, ?);`
	result, err := opr.db.Exec(insertoutgoingProductQuery, op.SoldAmount, op.SellingPrice, op.TotalSellingPrice, op.Notes, op.ProductID)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return err
	}

	op.ID = int(id)
	return err
}

func (opr *outgoingProductRepositoryImpl) AddWithTimestamps(op *model.OutgoingProduct) error {
	insertoutgoingProductQuery := `INSERT INTO OutgoingProducts (sold_amount, selling_price, total_selling_price, notes, product_id, created_at) VALUES( ?, ?, ?, ?, ?, ?);`
	result, err := opr.db.Exec(insertoutgoingProductQuery, op.SoldAmount, op.SellingPrice, op.TotalSellingPrice, op.Notes, op.ProductID, op.CreatedAt)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return err
	}

	op.ID = int(id)
	return err
}

func (opr *outgoingProductRepositoryImpl) GetSoldProducts(date model.Date) ([]model.SoldProduct, error) {
	var soldProducts []model.SoldProduct
	err := opr.db.Select(&soldProducts, "SELECT notes, created_at, SKU, name, sold_amount, selling_price FROM OutgoingProducts LEFT JOIN Products ON Products.SKU = OutgoingProducts.product_id WHERE created_at BETWEEN ? AND ?;", date.StartDate, date.EndDate)
	return soldProducts, err
}

func (opr *outgoingProductRepositoryImpl) GetTotalSoldProductByDate(date model.Date) (int, error) {
	var totalSoldProduct int
	err := opr.db.Get(&totalSoldProduct, "SELECT COUNT(DISTINCT(id)) FROM OutgoingProducts where created_at BETWEEN ? AND ? ;", date.StartDate, date.EndDate)
	return totalSoldProduct, err
}

func NewOutgoingProductRepository(Db *sqlx.DB) OutgoingProductRepository {
	return &outgoingProductRepositoryImpl{db: Db}
}
