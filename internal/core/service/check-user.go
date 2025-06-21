package service

import "github.com/NonthapatKim/kidzzle-api/internal/core/domain"

func (s *service) CheckUser(req domain.CheckUserRequest) (domain.CheckUserResult, error) {
	result, err := s.repo.CheckUser(req)
	if err != nil {
		return domain.CheckUserResult{}, err
	}

	return result, err
}
