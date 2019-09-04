# Description

Simple Dashboard for User to upload image and get the optimized image.

## Dev environment setup

#### 1. Golang

version : `1.12.5`

Install Golang with homebrew
```
brew install go
```

Or you can download Golang from (https://golang.org/dl/) and install it manually. 

#### 2. Sqlite3

version : `3.24.0`

Install Sqlite3 with homebrew
```
brew install sqlite3
```

Or you can download Golang from (https://golang.org/dl/) and install it manually. 

## Steps to run

You can run the file with run the binary
```
./ijahstore
```

Or ig you want to run manually you can do this steps

#### 1. Download packages


```
go mod download
```

#### 2. Migrate db
```
dbmate up
```

#### 3. Run go
```
go run main.go
```



## Endpoints

| Method  | Endpoint | Detail |
| ------------- | ------------- | ------------- | 
| GET | /products | Stores actual stock of products |  
| POST | /products| Get actual stock of products|
| POST | /incoming-products | To store product that will be stored into the inventory. |  
| GET | /incoming-products| Get history stored product |
| POST | /outgoing-products| To store product, quantity, notes of the products going out of inventory
|
| GET | /outgoing-products | Get history of product that going out of inventory |  
| GET | /reports | Shows a report for ijah to help her analyze and make decision. This report is related to total inventory value of Toko Ijah.|  
| GET | /sales-reports | Shows a report for ijah to help her analyze and make decision. This report is related to omzet / selling / profit.|
| POST | /migrate-data-products| Migrate data actual stock of products|
| POST | /migrate-data-incoming-products| To store product that will be stored into the inventory|
| POST | /migrate-data-outgoing-products| Migrate data to store product, quantity, notes of the products going out of inventory|



| Get| / | Open page which connected to websocket|

Raw body /products
```json
{
	"SKU" : "SSI-D01401071-LL-RED",
	"nama_barang" : "Zeomila Zipper Casual Blouse (L,Red)",
	"jumlah_barang" : 10
}
```


Raw body /incoming-products
```json
{
	"product_id" : "SSI-D01401071-LL-RED",
	"total_order" : 125,
	"total_received_order" : 125,
	"purchase_price": 68000,
	"notes" : "2017/05/18 terima 125",
	"receipt_number" : "(Hilang)"
}
```

Raw body /outgoing-products
```json
{
	"product_id" : "SSI-D01401071-LL-RED",
	"sold_amount" : 1,
	"selling_price" : 115000,
	"total_selling_price" : 115000,
	"notes" : "Pesanan ID-20180109-853724"
}
```

Example params on /sales-reports
```
/sales-reports?start_date=2017-08-11&end_date=2017-12-11
```

## Example Data

You can use csv files in example folder to migrate data
