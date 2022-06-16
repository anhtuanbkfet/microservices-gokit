package account

import (
	"context"
	"errors"
	"gokit-example/security"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gofrs/uuid"
)

type Service interface {
	CreateUser(ctx context.Context, email string, password string) (string, error)
	GetUser(ctx context.Context, id string) (string, error)
	ValidateUser(ctx context.Context, email, password string) (string, error)
	ValidateToken(ctx context.Context, token string) (string, error)
}

// ServiceMiddleware is a chainable behavior modifier for StringService.
type ServiceMiddleware func(Service) Service

var (
	ErrInvalidUser  = errors.New("invalid user")
	ErrInvalidToken = errors.New("invalid token")
)

/////////////////////////////////////////////////////////////////////////

type userService struct {
	repostory Repository
	logger    log.Logger
}

func NewService(rep Repository, logger log.Logger) Service {
	return &userService{
		repostory: rep,
		logger:    logger,
	}
}

func (s *userService) CreateUser(ctx context.Context, email string, password string) (string, error) {
	logger := log.With(s.logger, "method", "CreateUser")

	uuid, _ := uuid.NewV4()
	id := uuid.String()
	user := User{
		Id:       id,
		Email:    email,
		Password: password,
	}

	if err := s.repostory.CreateUser(ctx, user); err != nil {
		level.Error(logger).Log("error", err)
		return "", err
	}

	logger.Log("create user", id)

	return "Success", nil
}

func (s *userService) GetUser(ctx context.Context, id string) (string, error) {
	logger := log.With(s.logger, "method", "GetUser")

	email, err := s.repostory.GetUser(ctx, id)

	if err != nil {
		level.Error(logger).Log("error", err)
		return "", err
	}

	logger.Log("Get user", id)

	return email, nil
}

func (s *userService) ValidateUser(ctx context.Context, email, password string) (string, error) {
	//Check user and password:
	id, err := s.repostory.ValidateUser(ctx, email, password)

	if err != nil || id == "" {
		return "", ErrInvalidUser
	}
	token, err := security.NewToken(email)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *userService) ValidateToken(ctx context.Context, token string) (string, error) {
	t, err := security.ParseToken(token)
	if err != nil {
		return "", ErrInvalidToken
	}
	tData, err := security.GetClaims(t)
	if err != nil {
		return "", ErrInvalidToken
	}
	return tData["email"].(string), nil
}
