package service

import (
	"errors"

	"github.com/NonthapatKim/kidzzle-api/internal/core/domain"
	"github.com/NonthapatKim/kidzzle-api/internal/core/function"
	"gorm.io/gorm"
)

func (s *service) UserLogin(req domain.UserLoginRequest) (domain.UserLoginResponse, error) {
	var validationErrors domain.ValidationErrorResponse
	var response domain.UserLoginResponse
	var message string

	err := validate.Struct(req)
	if err != nil {
		validationErrors = ProcessValidationError(err)
		return response, domain.ValidationError{Errors: validationErrors}
	}

	loginResult, err := s.repo.UserLogin(req)
	if err != nil {
		if err.Error() == "incorrect password" {
			message = "อีเมลหรือรหัสผ่านไม่ถูกต้อง กรุณาลองใหม่อีกครั้ง"
			validationErrors.Incorrect = &message
			return response, domain.ValidationError{Errors: validationErrors}
		}

		if err.Error() == "user not found" {
			message = "ไม่พบบัญชีผู้ใช้นี้ กรุณาลองใหม่อีกครั้ง"
			validationErrors.UserError = &message
			return response, domain.ValidationError{Errors: validationErrors}
		}

		if err == gorm.ErrRecordNotFound {
			message = "อีเมลหรือรหัสผ่านไม่ถูกต้อง กรุณาลองใหม่อีกครั้ง"
			validationErrors.Incorrect = &message
			return response, domain.ValidationError{Errors: validationErrors}
		}
		return response, err
	}

	accessToken, err := function.GenerateAccessToken(loginResult.UserId)
	if err != nil {
		return response, errors.New("error generating token")
	}

	response = domain.UserLoginResponse{
		Code:        200,
		Message:     "successfully login",
		AccessToken: accessToken,
	}

	return response, nil
}
