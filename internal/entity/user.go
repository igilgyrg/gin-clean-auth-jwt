package entity

import (
	"context"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID ID `json:"id"`
	Email string `json:"email" binding:"email"`
	Password string `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

type UserUseCase interface {
	Get(ctx context.Context, user *User) (*User, error)
	GetByUsername(ctx context.Context, username string) (*User, error)
	Update(ctx context.Context, user *User) error
	Create(ctx context.Context, user *User) error
	Delete(ctx context.Context, id ID) error
}

type UserRepository interface{
	Get(id ID) (*User, error)
	GetByUsername(username string) (*User, error)
	Update(user User) error
	Save(user User) error
	Delete(id ID) error
}