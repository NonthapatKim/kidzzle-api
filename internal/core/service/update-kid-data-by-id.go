package service

import (
	"errors"

	"github.com/NonthapatKim/kidzzle-api/internal/core/domain"
	"github.com/NonthapatKim/kidzzle-api/internal/core/function"
)

func (s *service) UpdateKidDataById(req domain.UpdateKidDataByIdRequest) (domain.UpdateKidDataByIdResponse, error) {
	if req.AccessToken == "" {
		return domain.UpdateKidDataByIdResponse{}, errors.New("token is required")
	}

	if req.KidId == "" {
		return domain.UpdateKidDataByIdResponse{}, errors.New("kid_id is required")
	}

	if req.PregnantId == "" {
		return domain.UpdateKidDataByIdResponse{}, errors.New("pregnant_id is required")
	}

	_, err := function.ValidateAccessToken(&req.AccessToken)
	if err != nil {
		return domain.UpdateKidDataByIdResponse{}, err
	}

	_, err = s.repo.UpdateKidDataById(req)
	if err != nil {
		return domain.UpdateKidDataByIdResponse{}, err
	}

	response := domain.UpdateKidDataByIdResponse{
		Code:    200,
		Message: "Success",
	}

	return response, nil
}
