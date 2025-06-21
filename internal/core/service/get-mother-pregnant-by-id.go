package service

import (
	"errors"

	"github.com/NonthapatKim/kidzzle-api/internal/core/domain"
	"github.com/NonthapatKim/kidzzle-api/internal/core/function"
)

func (s *service) GetMotherPregnantById(req domain.GetMotherPregnantByIdRequest) ([]domain.GetMotherPregnantByIdResponse, error) {
	if req.AccessToken == "" {
		return nil, errors.New("token is required")
	}

	userId, err := function.ValidateAccessToken(&req.AccessToken)
	if err != nil {
		return nil, err
	}

	req.UserId = userId

	response, err := s.repo.GetMotherPregnantById(req)
	if err != nil {
		return nil, err
	}

	return response, nil
}
