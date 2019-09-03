package model

type Product struct {
	SKU   string `db:"SKU" json:"SKU"`
	Name  string `db:"name" json:"nama_barang"`
	Total int    `db:"total" json:"jumlah_barang"`
}
