package service

import (
	"errors"

	"github.com/NonthapatKim/kidzzle-api/internal/core/domain"
	"github.com/NonthapatKim/kidzzle-api/internal/core/function"
)

func (s *service) CreateAssessmentResult(req domain.CreateAssessmentResultRequest) (domain.CreateAssessmentResultResponse, error) {
	if req.AccessToken == "" {
		return domain.CreateAssessmentResultResponse{}, errors.New("token is required")
	}

	_, err := function.ValidateAccessToken(&req.AccessToken)
	if err != nil {
		return domain.CreateAssessmentResultResponse{}, err
	}

	_, err = s.repo.CreateAssessmentResult(req)
	if err != nil {
		return domain.CreateAssessmentResultResponse{}, err
	}

	response := domain.CreateAssessmentResultResponse{
		Code:    200,
		Message: "Success",
	}

	return response, nil
}
