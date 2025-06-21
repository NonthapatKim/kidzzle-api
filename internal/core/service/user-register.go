package service

import (
	"github.com/NonthapatKim/kidzzle-api/internal/core/domain"
	"github.com/NonthapatKim/kidzzle-api/internal/core/function"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var validate = validator.New()

func init() {
	validate.RegisterValidation("customEmail", function.CustomEmailValidator)
	validate.RegisterValidation("customPassword", function.CustomPasswordValidator)
}

func (s *service) UserRegister(req domain.UserRegisterRequest) (domain.UserRegisterResponse, error) {
	var validationErrors domain.ValidationErrorResponse
	var response domain.UserRegisterResponse
	var message string

	if req.UserId == uuid.Nil {
		req.UserId = uuid.New()
	}

	err := validate.Struct(req)
	if err != nil {
		validationErrors = ProcessValidationError(err)
		return response, domain.ValidationError{Errors: validationErrors}
	}

	reqCheckUser := domain.CheckUserRequest{
		Email: req.Email,
	}

	userExists, err := s.repo.CheckUser(reqCheckUser)
	if err != nil {
		return domain.UserRegisterResponse{}, err
	}

	if userExists.Exists {
		message = "user already exists"
		response = domain.UserRegisterResponse{
			Code:    409,
			Message: message,
		}
		return response, nil
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return response, err
	}
	hashedPasswordStr := string(hashedPassword)
	req.Password = hashedPasswordStr

	_, err = s.repo.UserRegister(req)
	if err != nil {
		return response, err
	}

	response = domain.UserRegisterResponse{
		Code:    200,
		Message: "successfully created user",
	}

	return response, nil
}
