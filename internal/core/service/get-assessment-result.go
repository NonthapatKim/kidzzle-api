package service

import (
	"errors"

	"github.com/NonthapatKim/kidzzle-api/internal/core/domain"
	"github.com/NonthapatKim/kidzzle-api/internal/core/function"
)

func (s *service) GetAssessmentResultById(req domain.GetAssessmentResultByIdRequest) ([]domain.GetAssessmentResultByIdResponse, error) {
	if req.AccessToken == "" {
		return nil, errors.New("token is required")
	}

	_, err := function.ValidateAccessToken(&req.AccessToken)
	if err != nil {
		return nil, err
	}

	result, err := s.repo.GetAssessmentResultById(req)
	if err != nil {
		return nil, err
	}

	return result, nil
}
