package service

import (
	"errors"

	"github.com/NonthapatKim/kidzzle-api/internal/core/domain"
	"github.com/NonthapatKim/kidzzle-api/internal/core/function"
)

func (s *service) DeleteMotherPregnantById(req domain.DeleteMotherPregnantByIdRequest) (domain.DeleteMotherPregnantByIdResponse, error) {
	if req.AccessToken == "" {
		return domain.DeleteMotherPregnantByIdResponse{}, errors.New("token is required")
	}

	_, err := function.ValidateAccessToken(&req.AccessToken)
	if err != nil {
		return domain.DeleteMotherPregnantByIdResponse{}, err
	}

	_, err = s.repo.DeleteMotherPregnantById(req)
	if err != nil {
		return domain.DeleteMotherPregnantByIdResponse{}, err
	}

	reponse := domain.DeleteMotherPregnantByIdResponse{
		Code:    200,
		Message: "success",
	}

	return reponse, nil
}
