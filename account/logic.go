package account

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/log/level"
	uuid "github.com/gofrs/uuid"
)

type service struct {
	repository Repository
	logger     log.Logger
}

func NewService(repo Repository, logger log.Logger) Service {
	return &service{
		repository: repo,
		logger:     logger,
	}
}

func (s service) CreateUser(ctx context.Context, email, password string) (string, error) {
	logger := log.With(s.logger, "method", "CreateUser")

	uuid, _ := uuid.NewV4()
	id := uuid.String()

	user := User{
		Id:       id,
		Email:    email,
		Password: password,
	}

	if err := s.repository.CreateUser(ctx, user); err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("Create User", user.ID)

	return user.ID, nil
}

func (s service) GetUser(ctx context.Context, id string) (string, error) {
	logger := log.With(s.logger, "method", "GetUser")

	email, err := s.repository.GetUser(ctx, id)

	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("Get user", id, "email", email)

	return email, nil
}
