package entities

import (
	"time"

	"github.com/google/uuid"
	usertype "github.com/joaofilippe/discount-club/domain/enums/user_type"
)

type User struct {
	id        uuid.UUID
	email     string
	password  string
	createdAt time.Time
	updatedAt time.Time
	deletedAt time.Time
	userType  usertype.UserType
	active    bool
}

func NewUser(email, password string, userType usertype.UserType) *User {
	return &User{
		email:    email,
		password: password,
		userType: userType,
		active:   true,
	}
}

func (u *User) ID() uuid.UUID {
	return u.id
}

func (u *User) Email() string {
	return u.email
}

func (u *User) Password() string {
	return u.password
}

func (u *User) CreatedAt() time.Time {
	return u.createdAt
}

func (u *User) UpdatedAt() time.Time {
	return u.updatedAt
}

func (u *User) DeletedAt() time.Time {
	return u.deletedAt
}

func (u *User) UserType() usertype.UserType {
	return u.userType
}

func (u *User) Active() bool {
	return u.active
}

func (u *User) SetID(id uuid.UUID) {
	u.id = id
}
