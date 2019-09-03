package model

import (
	"time"
)

type ReportProductValues struct {
	PrintDate    time.Time      `json:"printed_date"`
	TotalSKU     int            `json:"total_sku"`
	TotalProduct int            `json:"total_product"`
	TotalValue   int            `json:"total_value"`
	ProductValue []ProductValue `json:"product_value"`
}

type ProductValue struct {
	SKU              string `db:"SKU" json:"SKU"`
	Name             string `db:"name" json:"name"`
	TotalCurrentItem int    `db:"total" json:"total_current_item"`
	AveragePrice     int    `db:"average_price" json:"average_price"`
	TotalPrice       int    `db:"-" json:"total_price"`
}

func (pv *ProductValue) SetTotalPrice() {
	pv.TotalPrice = pv.TotalCurrentItem * pv.AveragePrice
}
