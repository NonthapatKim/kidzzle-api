package service

import "github.com/NonthapatKim/kidzzle-api/internal/core/domain"

func (s *service) GetPromotionColor() ([]domain.GetPromotionColorResponse, error) {
	result, err := s.repo.GetPromotionColor()
	if err != nil {
		return nil, err
	}

	return result, nil
}
