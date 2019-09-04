package service

import (
	"context"
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/adhistria/ijahstore/model"
	"github.com/adhistria/ijahstore/repository"
	"github.com/leekchan/accounting"
)

type (
	ReportService interface {
		GetProductValuesReports(context.Context) (string, error)
		WriteProductValuesReportToCSV(model.ReportProductValues) (string, error)
		GetSalesReports(context.Context, model.Date) (string, error)
		WriteSalesReportToCSV(model.SalesReport) (string, error)
	}

	reportServiceImpl struct {
		incomingProductRepository repository.IncomingProductRepository
		outgoingProductRepository repository.OutgoingProductRepository
	}
)

func (rs *reportServiceImpl) GetProductValuesReports(ctx context.Context) (string, error) {
	var productValues []model.ProductValue
	var reportProductValues model.ReportProductValues

	productValues, err := rs.incomingProductRepository.GetProductValues()

	if err != nil {
		return "", err
	}

	for index := range productValues {
		productValues[index].SetTotalPrice()
		reportProductValues.TotalProduct += productValues[index].TotalCurrentItem
		reportProductValues.TotalValue += productValues[index].TotalPrice
	}

	reportProductValues.PrintDate = time.Now()
	reportProductValues.TotalSKU = len(productValues)
	reportProductValues.ProductValue = productValues

	fileName, err := rs.WriteProductValuesReportToCSV(reportProductValues)

	return fileName, err
}

func (rs *reportServiceImpl) WriteProductValuesReportToCSV(data model.ReportProductValues) (string, error) {
	dateString := data.PrintDate
	dateString.String()
	dateString.Format("8 Jan 2006")
	fileName := dateString.Format("8 Jan 2006") + "result.csv"
	file, err := os.Create(fileName)
	if err != nil {
		log.Print("Error when create file")
		return fileName, err
	}
	writer := csv.NewWriter(file)

	ac := accounting.Accounting{Symbol: "Rp.", Precision: 0, Thousand: "."}

	writer.Write([]string{"Laporan Nilai Barang"})
	writer.Write([]string{""})
	writer.Write([]string{"Tanggal Cetak", dateString.Format("8 Jan 2006")})
	writer.Write([]string{"Jumlah SKU", strconv.Itoa(data.TotalSKU)})
	writer.Write([]string{"Jumlah Total Barang", strconv.Itoa(data.TotalProduct)})
	writer.Write([]string{"Total Nilai", ac.FormatMoney(data.TotalValue)})
	writer.Write([]string{""})
	writer.Write([]string{""})
	writer.Write([]string{""})
	writer.Write([]string{""})
	writer.Write([]string{"SKU", "Nama Item", "Jumlah", "Rata-Rata Harga Beli", "Total"})
	for _, productValue := range data.ProductValue {
		var strArray []string
		strArray = append(strArray, productValue.SKU)
		strArray = append(strArray, productValue.Name)
		strArray = append(strArray, strconv.Itoa(productValue.TotalCurrentItem))
		strArray = append(strArray, ac.FormatMoney(productValue.AveragePrice))
		strArray = append(strArray, ac.FormatMoney(productValue.TotalPrice))

		err := writer.Write(strArray)
		if err != nil {
			log.Print("error when write")
			return fileName, err
		}
	}

	writer.Flush()

	return fileName, err
}

func (rs *reportServiceImpl) GetSalesReports(ctx context.Context, date model.Date) (string, error) {

	var soldProducts []model.SoldProduct
	var salesReport model.SalesReport

	soldProducts, err := rs.outgoingProductRepository.GetSoldProducts(date)

	if err != nil {
		return "", err
	}

	for index, sp := range soldProducts {
		soldProducts[index].SetTotalSellingPrice()
		soldProducts[index].PurchasePrice, _ = rs.incomingProductRepository.GetAveragePriceByProductIDAndTimestamps(sp.SKU, sp.Date)
		soldProducts[index].SetProfit()
		salesReport.TotalOmzet += soldProducts[index].TotalSellingPrice
		salesReport.TotalGrossProfit += soldProducts[index].TotalSellingPrice
		salesReport.TotalProduct++
	}

	salesReport.PrintDate = date.StartDate + " - " + date.EndDate
	salesReport.TotalSales = len(soldProducts)
	salesReport.SoldProducts = soldProducts

	if err != nil {
		return "", err
	}

	fileName, err := rs.WriteSalesReportToCSV(salesReport)

	return fileName, err
}

func (rs *reportServiceImpl) WriteSalesReportToCSV(data model.SalesReport) (string, error) {
	fileName := data.PrintDate + "result.csv"
	file, err := os.Create(fileName)
	if err != nil {
		log.Print("Error when create file")
		return fileName, err
	}
	writer := csv.NewWriter(file)

	ac := accounting.Accounting{Symbol: "Rp.", Precision: 0, Thousand: "."}

	writer.Write([]string{"Laporan Penjualan"})
	writer.Write([]string{""})
	writer.Write([]string{"Tanggal Cetak", data.PrintDate})
	writer.Write([]string{"Laporan Omzet", ac.FormatMoney(data.TotalOmzet)})
	writer.Write([]string{"Laporan Laba Kotor", ac.FormatMoney(data.TotalGrossProfit)})
	writer.Write([]string{"Laporan Penjualan", strconv.Itoa(data.TotalSales)})
	writer.Write([]string{"Laporan Barang", strconv.Itoa(data.TotalProduct)})
	writer.Write([]string{""})
	writer.Write([]string{""})
	writer.Write([]string{""})
	writer.Write([]string{""})
	writer.Write([]string{"ID Pesanan", "Waktu", "SKU", "Nama Barang", "Jumlah", "Harga Jual", "Total", "Harga Beli", "Laba"})
	for _, soldProduct := range data.SoldProducts {
		var strArray []string

		strArray = append(strArray, soldProduct.Notes)
		strArray = append(strArray, soldProduct.Date)
		strArray = append(strArray, soldProduct.SKU)
		strArray = append(strArray, soldProduct.Name)
		strArray = append(strArray, strconv.Itoa(soldProduct.SoldAmount))
		strArray = append(strArray, ac.FormatMoney(soldProduct.SellingPrice))
		strArray = append(strArray, ac.FormatMoney(soldProduct.TotalSellingPrice))
		strArray = append(strArray, ac.FormatMoney(soldProduct.PurchasePrice))
		strArray = append(strArray, ac.FormatMoney(soldProduct.Profit))

		err := writer.Write(strArray)
		if err != nil {
			log.Print("error when write")
			return fileName, err
		}
	}

	writer.Flush()

	return fileName, err
}

func NewReportProductValueService(ipr repository.IncomingProductRepository, opr repository.OutgoingProductRepository) ReportService {
	return &reportServiceImpl{
		incomingProductRepository: ipr,
		outgoingProductRepository: opr,
	}
}
