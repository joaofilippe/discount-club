package requests

import (
	"fmt"

	usertype "github.com/joaofilippe/discount-club/app/domain/enums/user_type"
)

type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	UserType string `json:"user_type"`
}

func (r *CreateUserRequest) Parse() (*CreateUserRequest, error) {
	if r.Email == "" {
		return nil, fmt.Errorf("email is required")
	}
	if r.Password == "" {
		return nil, fmt.Errorf("password is required")
	}

	// Validate UserType
	parsed, ok := usertype.TryParse(r.UserType)
	if !ok {
		return nil, fmt.Errorf("invalid user type")
	}

	// Double check strict validation if needed, but TryParse covers valid enum values.
	// For strict matching against the subset of types allowed (if any restriction existed), we'd check here.
	// Since all types in Enum are valid:
	_ = parsed

	return r, nil
}
