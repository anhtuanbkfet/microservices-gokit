package account

import "context"

type User struct {
	Id       string `json: "id, omitempty"`
	Email    string `json: "email"`
	Password string `json: "password"`
}

type Repository interface {
	CreateUser(ctx context.Context, user User) error
	GetUser(ctx context.Context, id string) (string, error)
	ValidateUser(ctx context.Context, email, password string) (string, error)
}
