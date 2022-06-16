package account

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type (
	CreateUserRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	CreateUserResponse struct {
		Ok string `json:"ok"`
	}

	GetUserRequest struct {
		Id string `json:"id"`
	}
	GetUserResponse struct {
		Email string `json:"email"`
	}

	ValidateUserRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	ValidateUserResponse struct {
		Token string `json:"token,omitempty"`
		Err   string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
	}

	ValidateTokenRequest struct {
		Token string `json:"token"`
	}

	ValidateTokenResponse struct {
		Email string `json:"email,omitempty"`
		Err   string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
	}
)

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeEmailRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req GetUserRequest
	vars := mux.Vars(r)

	req = GetUserRequest{
		Id: vars["id"],
	}
	return req, nil
}

func decodeValidateUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request ValidateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeValidateTokenRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request ValidateTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
