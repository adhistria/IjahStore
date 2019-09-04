package model

type SalesReport struct {
	PrintDate string `json:"printed_date"`
	// Date             time.Time     `json:"date"`
	TotalOmzet       int           `json:"total_omzet"`
	TotalGrossProfit int           `json:"total_gross_profit"`
	TotalSales       int           `json:"total_sales"`
	TotalProduct     int           `json:"total_product"`
	SoldProducts     []SoldProduct `json:"sold_products"`
}

type SoldProduct struct {
	Notes             string `db:"notes" `
	Date              string `db:"created_at" `
	SKU               string `db:"SKU" `
	Name              string `db:"name" `
	SoldAmount        int    `db:"sold_amount" `
	SellingPrice      int    `db:"selling_price" `
	TotalSellingPrice int    `db:"-" `
	PurchasePrice     int    `db:"-" `
	Profit            int    `db:"-" `
}

type Date struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

func (ps *SoldProduct) SetTotalSellingPrice() {
	ps.TotalSellingPrice = ps.SellingPrice * ps.SoldAmount
}

func (ps *SoldProduct) SetProfit() {
	ps.Profit = ps.SellingPrice - ps.PurchasePrice
}
