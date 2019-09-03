package controller

import (
	"io"
	"net/http"
	"os"
	"strconv"

	"context"
	"log"

	"github.com/adhistria/ijahstore/response"
	"github.com/adhistria/ijahstore/service"
	"github.com/gorilla/mux"
)

type reportController struct {
	ReportService service.ReportProductValueService
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

func NewReportController(router *mux.Router, ps service.ReportProductValueService) reportController {
	controller := reportController{ReportService: ps}
	router.HandleFunc("/reports", controller.GetProductValuesReports).Methods("GET")
	return controller
}
