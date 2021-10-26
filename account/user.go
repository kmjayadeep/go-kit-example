package account

import "context"

type User struct {
	ID       string `json:"id,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type Repository interface {
	CreateUser(ctx context.Context, userp User) error
	GetUser(ctx context.Context, id string) (string, error)
}
