package controller

import (
	"io"
	"net/http"
	"os"
	"strconv"

	"context"
	"log"

	"github.com/adhistria/ijahstore/model"
	"github.com/adhistria/ijahstore/response"
	"github.com/adhistria/ijahstore/service"
	"github.com/gorilla/mux"
)

type reportController struct {
	ReportService service.ReportService
}

func (rc *reportController) GetProductValuesReports(w http.ResponseWriter, r *http.Request) {

	fileName, err := rc.ReportService.GetProductValuesReports(context.Background())

	if err != nil {
		log.Print(err.Error())
		response.APIErrorResponse(w, 500, err.Error())
		return
	}

	reportFile, err := os.Open(fileName)
	defer reportFile.Close()

	if err != nil {
		log.Print(err.Error())
		response.APIErrorResponse(w, 500, err.Error())
		return
	}

	fileHeader := make([]byte, 512)
	reportFile.Read(fileHeader)
	fileContentType := http.DetectContentType(fileHeader)

	fileStat, _ := reportFile.Stat()
	fileSize := strconv.FormatInt(fileStat.Size(), 10)

	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Type", fileContentType)
	w.Header().Set("Content-Length", fileSize)

	reportFile.Seek(0, 0)
	io.Copy(w, reportFile)
	return
}

func (rc *reportController) GetSalesReports(w http.ResponseWriter, r *http.Request) {

	startDateParams, ok := r.URL.Query()["start_date"]
	if !ok || len(startDateParams[0]) < 1 {
		response.APIErrorResponse(w, 500, "Url Param 'start_date' is missing")
		return
	}

	endDateParams, ok := r.URL.Query()["end_date"]
	if !ok || len(endDateParams[0]) < 1 {
		response.APIErrorResponse(w, 500, "Url Param 'end_date' is missing")
		return
	}

	startDate := startDateParams[0]
	endDate := endDateParams[0]

	log.Println("Url Param date is: " + string(startDate) + " " + string(endDate))

	var date model.Date
	date.StartDate = startDate
	date.EndDate = endDate

	fileName, err := rc.ReportService.GetSalesReports(context.Background(), date)

	reportFile, err := os.Open(fileName)
	defer reportFile.Close()

	if err != nil {
		response.APIErrorResponse(w, 500, err.Error())
		return
	}

	fileHeader := make([]byte, 512)
	reportFile.Read(fileHeader)
	fileContentType := http.DetectContentType(fileHeader)

	fileStat, _ := reportFile.Stat()
	fileSize := strconv.FormatInt(fileStat.Size(), 10)

	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Type", fileContentType)
	w.Header().Set("Content-Length", fileSize)

	reportFile.Seek(0, 0)
	io.Copy(w, reportFile)
	return

}

func NewReportController(router *mux.Router, ps service.ReportService) reportController {
	controller := reportController{ReportService: ps}
	router.HandleFunc("/reports", controller.GetProductValuesReports).Methods("GET")
	router.HandleFunc("/sales-reports", controller.GetSalesReports).Methods("GET")
	return controller
}
