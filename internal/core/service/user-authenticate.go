package service

import (
	"errors"

	"github.com/NonthapatKim/kidzzle-api/internal/core/domain"
	"github.com/NonthapatKim/kidzzle-api/internal/core/function"
)

func (s *service) UserAuthenticate(req domain.UserAuthenticateRequest) (domain.UserAuthenticateResponse, error) {
	if req.AccessToken == "" {
		return domain.UserAuthenticateResponse{}, errors.New("token is required")
	}

	userId, err := function.ValidateAccessToken(&req.AccessToken)
	if err != nil {
		return domain.UserAuthenticateResponse{}, err
	}

	response := domain.UserAuthenticateResponse{
		UserId: userId,
	}

	return response, nil
}
