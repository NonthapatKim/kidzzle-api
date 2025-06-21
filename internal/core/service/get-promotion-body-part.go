package service

import "github.com/NonthapatKim/kidzzle-api/internal/core/domain"

func (s *service) GetPromotionBodyPart() ([]domain.GetPromotionBodyPartResponse, error) {
	result, err := s.repo.GetPromotionBodyPart()
	if err != nil {
		return nil, err
	}

	return result, nil
}
