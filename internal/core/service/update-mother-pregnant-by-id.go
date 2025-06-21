package service

import (
	"errors"

	"github.com/NonthapatKim/kidzzle-api/internal/core/domain"
	"github.com/NonthapatKim/kidzzle-api/internal/core/function"
)

func (s *service) UpdateMotherPregnantById(req domain.UpdateMotherPregnantByIdRequest) (domain.UpdateMotherPregnantByIdResponse, error) {
	if req.AccessToken == "" {
		return domain.UpdateMotherPregnantByIdResponse{}, errors.New("token is required")
	}

	_, err := function.ValidateAccessToken(&req.AccessToken)
	if err != nil {
		return domain.UpdateMotherPregnantByIdResponse{}, err
	}

	_, err = s.repo.UpdateMotherPregnantById(req)
	if err != nil {
		return domain.UpdateMotherPregnantByIdResponse{}, err
	}

	response := domain.UpdateMotherPregnantByIdResponse{
		Code:    200,
		Message: "Success",
	}

	return response, nil
}
