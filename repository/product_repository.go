package repository

import (
	"github.com/adhistria/ijahstore/model"
	"github.com/jmoiron/sqlx"
)

type ProductRepository interface {
	Add(*model.Product) error
	FindAll() ([]model.Product, error)
	AddTotalProduct(*model.IncomingProduct) error
	GetProduct(string) (model.Product, error)
}

type productRepositoryImpl struct {
	db *sqlx.DB
}

func (pr *productRepositoryImpl) Add(p *model.Product) error {

	insertProductQuery := `INSERT OR REPLACE INTO Products (SKU, name, total) VALUES( ?, ?, ? ) ;`
	_, err := pr.db.Exec(insertProductQuery, p.SKU, p.Name, p.Total, p.Name, p.Total)

	return err
}

func (pr *productRepositoryImpl) FindAll() ([]model.Product, error) {
	var products []model.Product
	err := pr.db.Select(&products, "SELECT * FROM Products")
	return products, err
}

func (pr *productRepositoryImpl) AddTotalProduct(ip *model.IncomingProduct) error {
	updateProductQuery := `UPDATE Products SET total = total + ? WHERE SKU = ?`

	_, err := pr.db.Exec(updateProductQuery, ip.TotalReceiveOrder, ip.ProductID)

	return err

}

func (pr *productRepositoryImpl) GetProduct(productId string) (model.Product, error) {
	var product model.Product
	err := pr.db.Get(&product, "SELECT * FROM Products WHERE SKU = $1", productId)

	return product, err
}

func NewProductRepository(Db *sqlx.DB) ProductRepository {
	return &productRepositoryImpl{db: Db}
}
