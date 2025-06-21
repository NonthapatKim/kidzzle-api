package service

import (
	"errors"

	"github.com/NonthapatKim/kidzzle-api/internal/core/domain"
	"github.com/NonthapatKim/kidzzle-api/internal/core/function"
)

func (s *service) GetKidData(req domain.GetKidDataRequest) ([]domain.GetKidDataResponse, error) {
	if req.AccessToken == "" {
		return nil, errors.New("token is required")
	}

	if req.PregnantId == "" {
		return nil, errors.New("pregnant_id is required")
	}

	userId, err := function.ValidateAccessToken(&req.AccessToken)
	if err != nil {
		return nil, err
	}

	req.UserId = userId

	result, err := s.repo.GetKidData(req)
	if err != nil {
		return nil, err
	}

	return result, nil
}
