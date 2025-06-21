package service

import "github.com/NonthapatKim/kidzzle-api/internal/core/domain"

func (s *service) GetAssessmentTraining(req domain.GetAssessmentTrainingMethodsRequest) ([]domain.GetAssessmentTrainingMethodsResponse, error) {
	result, err := s.repo.GetAssessmentTraining(req)
	if err != nil {
		return nil, err
	}

	return result, nil
}
