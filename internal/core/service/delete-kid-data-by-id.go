package service

import (
	"errors"

	"github.com/NonthapatKim/kidzzle-api/internal/core/domain"
	"github.com/NonthapatKim/kidzzle-api/internal/core/function"
)

func (s *service) DeleteKidDataById(req domain.DeleteKidDataByIdRequest) (domain.DeleteKidDataByIdResponse, error) {
	if req.AccessToken == "" {
		return domain.DeleteKidDataByIdResponse{}, errors.New("token is required")
	}

	_, err := function.ValidateAccessToken(&req.AccessToken)
	if err != nil {
		return domain.DeleteKidDataByIdResponse{}, err
	}

	_, err = s.repo.DeleteKidDataById(req)
	if err != nil {
		return domain.DeleteKidDataByIdResponse{}, err
	}

	reponse := domain.DeleteKidDataByIdResponse{
		Code:    200,
		Message: "success",
	}

	return reponse, nil
}
