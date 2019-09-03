package model

type SuccessResponse struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}
