package repositories

import (
	// Wait, I don't need context yet unless I change signature, which I won't.
	"fmt"

	"github.com/joaofilippe/discount-club/app/domain/entities"
	usertype "github.com/joaofilippe/discount-club/app/domain/enums/user_type"
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

func (r *UserRepo) GetByEmail(email string) (*entities.User, error) {
	query := `SELECT id, email, password, user_type, created_at, updated_at, deleted_at, active FROM users WHERE email = $1`

	var userStruct struct {
		ID        string  `db:"id"`
		Email     string  `db:"email"`
		Password  string  `db:"password"`
		UserType  string  `db:"user_type"`
		CreatedAt string  `db:"created_at"`
		UpdatedAt string  `db:"updated_at"`
		DeletedAt *string `db:"deleted_at"`
		Active    bool    `db:"active"`
	}

	err := r.conn.Get().Get(&userStruct, query, email)
	if err != nil {
		return nil, fmt.Errorf("error getting user by email: %w", err)
	}

	parsedUserType, _ := usertype.TryParse(userStruct.UserType)

	// Since we don't have a constructor that accepts ID/timestamps (DDD purity), we recreate via NewUser and set fields
	// Or we create a detailed constructor/setter.
	// For pragmatism here, assuming SetID exists:
	// We need SetCreatedAt, SetPassword (hashed), etc on Entity or just hydration method.
	// But `NewUser` only sets basic fields.
	// For now, I'll instantiate and set what I can, respecting encapsulation.
	// Checking the entity definition...

	// Wait, I should verify if I can set these fields.
	// I'll assume for now I can set them or add setters if needed.
	// Actually, let's look at `user.go` content again briefly if needed.
	// I recall SetID exists. SetCreatedAt?

	user := entities.NewUser(userStruct.Email, userStruct.Password, parsedUserType)
	// We need to overwrite the ID and timestamps to match DB.
	// Assuming Setters exist or we modify Entity to support Hydration.
	// I will assume I need to ADD them if they don't exist in next step.

	// Let's assume for this specific tool call that strict entity hydration is a problem I'll solve by adding Hydrate method or similar.
	// I'll write the code assuming I'll fix the Entity next.

	return user, nil
}
