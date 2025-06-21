package service

import "github.com/NonthapatKim/kidzzle-api/internal/core/domain"

func (s *service) GetAssessmentQuestion(req domain.GetAssessmentQuestionRequest) ([]domain.GetAssessmentQuestionResponse, error) {
	result, err := s.repo.GetAssessmentQuestion(req)
	if err != nil {
		return nil, err
	}

	return result, nil
}
