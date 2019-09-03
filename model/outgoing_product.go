package model

import (
	"time"
)

type OutgoingProduct struct {
	ID                int       `db:"id" json:"id"`
	CreatedAt         time.Time `db:"created_at" json:"created_at"`
	ProductID         string    `db:"product_id" json:"product_id"`
	SoldAmount        int       `db:"sold_amount" json:"sold_amount"`
	SellingPrice      int       `db:"selling_price" json:"selling_price"`
	TotalSellingPrice int       `db:"total_selling_price" json:"total_selling_price"`
	Notes             string    `db:"notes" json:"notes"`
	Product           Product   `json:"product"`
}

func (op *OutgoingProduct) SetTotalSellingPrice() {
	op.TotalSellingPrice = op.SoldAmount * op.SellingPrice
}
