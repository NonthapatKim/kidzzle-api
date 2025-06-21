package service

import (
	"errors"

	"github.com/NonthapatKim/kidzzle-api/internal/core/domain"
	"github.com/NonthapatKim/kidzzle-api/internal/core/function"
)

func (s *service) CreateMotherPregnant(req domain.CreateMotherPregnantRequest) (domain.CreateMotherPregnantResponse, error) {
	if req.AccessToken == "" {
		return domain.CreateMotherPregnantResponse{}, errors.New("token is required")
	}

	userId, err := function.ValidateAccessToken(&req.AccessToken)
	if err != nil {
		return domain.CreateMotherPregnantResponse{}, err
	}

	req.UserId = userId

	_, err = s.repo.CreateMotherPregnant(req)
	if err != nil {
		return domain.CreateMotherPregnantResponse{}, err
	}

	response := domain.CreateMotherPregnantResponse{
		Code:    200,
		Message: "Success",
	}

	return response, nil
}
