package requests

import "fmt"

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *LoginRequest) Parse() (*LoginRequest, error) {
	if r.Email == "" {
		return nil, fmt.Errorf("email is required")
	}
	if r.Password == "" {
		return nil, fmt.Errorf("password is required")
	}
	return r, nil
}
