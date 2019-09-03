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
	ReportProductValueService interface {
		GetProductValuesReports(context.Context) (string, error)
		WriteProductValuesReportToCSV(model.ReportProductValues) (string, error)
	}

	reportProductValueServiceImpl struct {
		incomingProductRepository repository.IncomingProductRepository
		outgoingProductRepository repository.OutgoingProductRepository
	}
)

func (rs *reportProductValueServiceImpl) GetProductValuesReports(ctx context.Context) (string, error) {
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

func (rs *reportProductValueServiceImpl) WriteProductValuesReportToCSV(data model.ReportProductValues) (string, error) {
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

func NewReportProductValueService(ipr repository.IncomingProductRepository, opr repository.OutgoingProductRepository) ReportProductValueService {
	return &reportProductValueServiceImpl{
		incomingProductRepository: ipr,
		outgoingProductRepository: opr,
	}
}
