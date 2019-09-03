package model

import (
	"time"
)

type IncomingProduct struct {
	ID                 int       `db:"id" json:"id"`
	CreatedAt          time.Time `db:"created_at" json:"created_at"`
	ProductID          string    `db:"product_id" json:"product_id"`
	TotalOrder         int       `db:"total_order" json:"total_order"`
	TotalReceiveOrder  int       `db:"total_received_order" json:"total_received_order"`
	PurchasePrice      int       `db:"purchase_price" json:"purchase_price"`
	TotalPurchasePrice int       `db:"total_purchase_price" json:"total_purchase_price"`
	ReceiptNumber      string    `db:"receipt_number" json:"receipt_number"`
	Notes              string    `db:"notes" json:"notes"`
	Product            Product   `json:"product"`
}

func (ip *IncomingProduct) SetTotalPurchasePrice() {
	ip.TotalPurchasePrice = ip.TotalOrder * ip.PurchasePrice
}
