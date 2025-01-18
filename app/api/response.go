package api

type CommonResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type CommonErrorResponse struct {
	Code         int    `json:"code"`
	ErrorMessage string `json:"error_message"`
	ErrorDetails any    `json:"error_details"`
}
