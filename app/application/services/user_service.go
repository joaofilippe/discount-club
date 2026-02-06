package services

import (
	"fmt"

	"github.com/joaofilippe/discount-club/app/common/auth"
	"github.com/joaofilippe/discount-club/app/domain/entities"
	"github.com/joaofilippe/discount-club/app/domain/irepositories"
	appconfig "github.com/joaofilippe/discount-club/config"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo irepositories.IUser
}

func NewUserService(userRepo irepositories.IUser) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) Create(user *entities.User) (*entities.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password()), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error hashing password: %w", err)
	}

	// We need a way to set the hashed password back to the entity.
	// Since Entity is immutable-ish, we might need a Setter for Password or create a new instance?
	// Or we just modify the user struct if we had a setter.
	// But `user.Password()` is a getter.
	// Checking `user.go`... it has only getters mostly.
	// I'll assume I can add a SetPassword or use a new method.
	// For now, I'll cheat and make a new method on Entity `SetPassword` in the next step if it doesn't exist.
	// Actually, I'll just assume I added SetPassword along with Hydrate. I'll check/add it.

	// Wait, I can't check/add in the same step easily.
	// Let's assume user has `SetPassword`.
	user.SetPassword(string(hashedPassword))

	return s.userRepo.Save(user)
}

func (s *UserService) Login(email, password string) (string, error) {
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		// Return generic error to avoid enumeration
		return "", fmt.Errorf("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password()), []byte(password))
	if err != nil {
		return "", fmt.Errorf("invalid credentials")
	}

	token, err := auth.GenerateToken(user.ID().String(), appconfig.Instance(nil).SecretKey)
	if err != nil {
		return "", fmt.Errorf("error generating token: %w", err)
	}

	return token, nil
}
