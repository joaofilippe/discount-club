package repositories

import (
	"fmt"

	"github.com/joaofilippe/discount-club/app/domain/entities"
	"github.com/joaofilippe/discount-club/app/infra/database"
)

type UserRepo struct {
	conn *database.Connection
}

func NewUserRepo(conn *database.Connection) *UserRepo {
	return &UserRepo{
		conn,
	}
}

func (r *UserRepo) Save(user *entities.User) (*entities.User, error) {
	query := `INSERT INTO users (id, email, password, user_type, created_at, updated_at, deleted_at, active) 
		VALUES (:id, :email, :password, :user_type, :created_at, :updated_at, :deleted_at, :active)`

	newUser := struct {
		ID        string  `db:"id"`
		Email     string  `db:"email"`
		Password  string  `db:"password"`
		UserType  string  `db:"user_type"`
		CreatedAt string  `db:"created_at"`
		UpdatedAt string  `db:"updated_at"`
		DeletedAt *string `db:"deleted_at"`
		Active    bool    `db:"active"`
	}{
		ID:        user.ID().String(),
		Email:     user.Email(),
		Password:  user.Password(),
		UserType:  user.UserType().DBString(),
		CreatedAt: user.CreatedAt().Format("2006-01-02 15:04:05"),
		UpdatedAt: user.UpdatedAt().Format("2006-01-02 15:04:05"),
		Active:    user.Active(),
	}

	_, err := r.conn.Get().NamedExec(query, newUser)
	if err != nil {
		return nil, fmt.Errorf("error saving user: %w", err)
	}

	return user, nil
}
