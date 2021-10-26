package account

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"
	"github.com/pkg/errors"
)

type repo struct {
	store  map[string]User
	logger log.Logger
}

var ErrUserNotFound = errors.New("User not found")
var RepoErr = errors.New("Email or password empty")

func NewRepo(logger log.Logger) Repository {
	return &repo{
		store:  make(map[string]User),
		logger: logger,
	}
}

func (r *repo) CreateUser(ctx context.Context, user User) error {
	logger := log.With(r.logger, "method", "repo CreateUser")

	if user.Email == "" || user.Password == "" {
		return RepoErr
	}

	r.store[user.ID] = user
	logger.Log("store after saving", fmt.Sprint(r.store))
	return nil
}

func (r *repo) GetUser(ctx context.Context, id string) (string, error) {
	logger := log.With(r.logger, "method", "repo GetUser")

	user, ok := r.store[id]

	if !ok {
		return "", ErrUserNotFound
	}

	logger.Log("store after saving", r.store)
	return user.Email, nil
}
