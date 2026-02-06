package responses

import (
	"net/http"

	commonapi "github.com/joaofilippe/discount-club/app/api/common"
)

type LoginResponseDTO struct {
	Token string `json:"token"`
}

func BuildLoginResponse(token string) commonapi.CommonResponse {
	return commonapi.CommonResponse{
		Code:    http.StatusOK,
		Message: "Login successful",
		Data: LoginResponseDTO{
			Token: token,
		},
	}
}
