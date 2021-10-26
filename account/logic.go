package account

import (
  "context"

  "github.com/go-kit/kit/log"
)

type service struct {
  repository Repository
  logger log.Logger
}

func NewService(repo Repository, logger log.Logger) Service {
  return &service{
    repository: repo,
    logger: logger,
  }
}

func (s service) CreateUser(ctx context.Context, email, password string) (string, error) {
  // TODO complete logic
  return "", nil
}

func (s service) GetUser(ctx context.Context, id string) (string, error) {
  // TODO complete logic
  return "", nil
}
