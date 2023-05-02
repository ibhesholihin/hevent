package models

type (
	HttpResponse struct {
		Code    uint        `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}

	HTTPResponseWithoutData struct {
		Code    uint   `json:"code"`
		Message string `json:"message"`
	}
)
