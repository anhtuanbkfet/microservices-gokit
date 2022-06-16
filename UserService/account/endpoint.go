package account

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateUser    endpoint.Endpoint
	GetUser       endpoint.Endpoint
	ValidateUser  endpoint.Endpoint
	ValidateToken endpoint.Endpoint
}

func MakeEndpoints(srv Service) Endpoints {
	return Endpoints{
		CreateUser:    makeCreateUserEndpoint(srv),
		GetUser:       makeGetUserEndpoint(srv),
		ValidateUser:  makeValidateUserEndpoint(srv),
		ValidateToken: makeValidateTokenEndpoint(srv),
	}
}

func makeCreateUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateUserRequest)
		ok, err := s.CreateUser(ctx, req.Email, req.Password)
		return CreateUserResponse{Ok: ok}, err
	}
}

func makeGetUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserRequest)
		email, err := s.GetUser(ctx, req.Id)

		return GetUserResponse{
			Email: email,
		}, err
	}
}

func makeValidateUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ValidateUserRequest)
		token, err := s.ValidateUser(ctx, req.Email, req.Password)
		if err != nil {
			return ValidateUserResponse{"", err.Error()}, err
		}
		return ValidateUserResponse{token, ""}, err
	}
}

func makeValidateTokenEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ValidateTokenRequest)
		email, err := s.ValidateToken(ctx, req.Token)
		if err != nil {
			return ValidateTokenResponse{"", err.Error()}, err
		}
		return ValidateTokenResponse{email, ""}, err
	}
}
