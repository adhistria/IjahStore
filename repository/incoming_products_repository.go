package repository

import (
	"github.com/adhistria/ijahstore/model"
	"github.com/jmoiron/sqlx"
)

type IncomingProductRepository interface {
	Add(*model.IncomingProduct) error
	FindAll() ([]model.IncomingProduct, error)
	GetProductValues() ([]model.ProductValue, error)
	GetAveragePriceByProductIDAndTimestamps(productID string, createdAt string) (int, error)
	AddWithTimestamps(*model.IncomingProduct) error
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

func (ipr *incomingProductRepositoryImpl) GetProductValues() ([]model.ProductValue, error) {
	getProductValuesQuery := `SELECT (SUM(total_purchase_price) / SUM(total_received_order)) AS average_price,  SKU, total, name FROM IncomingProducts  LEFT JOIN Products ON product_id = SKU GROUP BY (product_id)`
	var productValues []model.ProductValue
	err := ipr.db.Select(&productValues, getProductValuesQuery)
	return productValues, err
}

func (ipr *incomingProductRepositoryImpl) GetAveragePriceByProductIDAndTimestamps(productID string, createdAt string) (int, error) {
	getAveragePricesQuery := `SELECT (SUM(total_purchase_price) / SUM(total_received_order)) AS average_price  FROM IncomingProducts  WHERE IncomingProducts.product_id = ? AND created_at <= ? GROUP BY (product_id) ;`
	var averagePrice int
	err := ipr.db.Get(&averagePrice, getAveragePricesQuery, productID, createdAt)
	return averagePrice, err
}

func (ipr *incomingProductRepositoryImpl) AddWithTimestamps(ip *model.IncomingProduct) error {
	insertIncomingProductQuery := `INSERT INTO IncomingProducts (total_order, total_received_order, purchase_price, total_purchase_price, receipt_number,  notes, product_id, created_at) VALUES(?, ?, ?, ?, ?, ?, ?, ?);`
	result, err := ipr.db.Exec(insertIncomingProductQuery, ip.TotalOrder, ip.TotalReceiveOrder, ip.PurchasePrice, ip.TotalPurchasePrice, ip.ReceiptNumber, ip.Notes, ip.ProductID, ip.CreatedAt)

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

func NewIncomingProductRepository(Db *sqlx.DB) IncomingProductRepository {
	return &incomingProductRepositoryImpl{db: Db}
}
