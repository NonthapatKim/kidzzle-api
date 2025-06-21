package service

import (
	"errors"

	"github.com/NonthapatKim/kidzzle-api/internal/core/domain"
	"github.com/NonthapatKim/kidzzle-api/internal/core/function"
)

func (s *service) CreateKidData(req domain.CreateKidDataRequest) (domain.CreateKidDataResponse, error) {
	if req.AccessToken == "" {
		return domain.CreateKidDataResponse{}, errors.New("token is required")
	}

	if req.PregnantId == "" {
		return domain.CreateKidDataResponse{}, errors.New("pregnant_id is required")
	}

	userId, err := function.ValidateAccessToken(&req.AccessToken)
	if err != nil {
		return domain.CreateKidDataResponse{}, err
	}

	req.UserId = userId

	_, err = s.repo.CreateKidData(req)
	if err != nil {
		return domain.CreateKidDataResponse{}, err
	}

	response := domain.CreateKidDataResponse{
		Code:    200,
		Message: "Success",
	}

	return response, nil
}
