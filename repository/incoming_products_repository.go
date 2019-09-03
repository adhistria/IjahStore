package repository

import (
	"github.com/adhistria/ijahstore/model"
	"github.com/jmoiron/sqlx"
)

type IncomingProductRepository interface {
	Add(*model.IncomingProduct) error
	FindAll() ([]model.IncomingProduct, error)
}

type incomingProductRepositoryImpl struct {
	db *sqlx.DB
}

func (ipr *incomingProductRepositoryImpl) Add(ip *model.IncomingProduct) error {
	insertIncomingProductQuery := `INSERT INTO IncomingProducts (total_order, total_received_order, purchase_price, total_purchase_price, receipt_number,  notes, product_id) VALUES(?, ?, ?, ?, ?, ?, ?);`
	result, err := ipr.db.Exec(insertIncomingProductQuery, ip.TotalOrder, ip.TotalReceiveOrder, ip.PurchasePrice, ip.TotalPurchasePrice, ip.ReceiptNumber, ip.Notes, ip.ProductID)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return err
	}

	ip.ID = int(id)
	return err

}

func (ipr *incomingProductRepositoryImpl) FindAll() ([]model.IncomingProduct, error) {
	var incomingProducts []model.IncomingProduct
	err := ipr.db.Select(&incomingProducts, "SELECT * FROM IncomingProducts")
	return incomingProducts, err
}

func NewIncomingProductRepository(Db *sqlx.DB) IncomingProductRepository {
	return &incomingProductRepositoryImpl{db: Db}
}
