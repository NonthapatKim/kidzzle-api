package port

import (
	"context"

	"github.com/NonthapatKim/kidzzle-api/internal/core/domain"
)

type Service interface {
	// User
	CheckUser(req domain.CheckUserRequest) (domain.CheckUserResult, error)
	UserAuthenticate(req domain.UserAuthenticateRequest) (domain.UserAuthenticateResponse, error)
	UserLogin(req domain.UserLoginRequest) (domain.UserLoginResponse, error)
	UserLoginBySocial(req domain.UserLoginBySocialRequest) (domain.UserLoginBySocialResponse, error)
	UserRegister(req domain.UserRegisterRequest) (domain.UserRegisterResponse, error)
	UserRegisterBySocial(req domain.UserRegisterBySocialRequest) (domain.UserRegisterBySocialResponse, error)
	UserRequestResetPassword(req domain.UserRequestResetPasswordRequest) (domain.UserRequestResetPasswordResponse, error)
	UserResetPassword(req domain.UserResetPasswordRequest) (domain.UserResetPasswordResponse, error)

	// Assessment
	CreateAssessmentResult(req domain.CreateAssessmentResultRequest) (domain.CreateAssessmentResultResponse, error)
	GetAssessmentAgeRange(req domain.GetAssessmentAgeRangeRequest) ([]domain.GetAssessmentAgeRangeResponse, error)
	GetAssessmentQuestion(req domain.GetAssessmentQuestionRequest) ([]domain.GetAssessmentQuestionResponse, error)
	GetAssessmentResultById(req domain.GetAssessmentResultByIdRequest) ([]domain.GetAssessmentResultByIdResponse, error)
	GetAssessmentResultDataById(req domain.GetAssessmentResultDataByIdRequest) ([]domain.GetAssessmentResultDataByIdResponse, error)
	GetAssessmentTraining(req domain.GetAssessmentTrainingMethodsRequest) ([]domain.GetAssessmentTrainingMethodsResponse, error)

	// Mother Pregnant
	CreateMotherPregnant(req domain.CreateMotherPregnantRequest) (domain.CreateMotherPregnantResponse, error)
	DeleteMotherPregnantById(req domain.DeleteMotherPregnantByIdRequest) (domain.DeleteMotherPregnantByIdResponse, error)
	GetMotherPregnantById(req domain.GetMotherPregnantByIdRequest) ([]domain.GetMotherPregnantByIdResponse, error)
	UpdateMotherPregnantById(req domain.UpdateMotherPregnantByIdRequest) (domain.UpdateMotherPregnantByIdResponse, error)

	// Kid
	CreateKidData(req domain.CreateKidDataRequest) (domain.CreateKidDataResponse, error)
	DeleteKidDataById(req domain.DeleteKidDataByIdRequest) (domain.DeleteKidDataByIdResponse, error)
	GetKidData(req domain.GetKidDataRequest) ([]domain.GetKidDataResponse, error)
	UpdateKidDataById(req domain.UpdateKidDataByIdRequest) (domain.UpdateKidDataByIdResponse, error)

	// Promotion
	GetPromotionBodyPart() ([]domain.GetPromotionBodyPartResponse, error)
	GetPromotionColor() ([]domain.GetPromotionColorResponse, error)

	// Other
	UserAuthByGoogle(ctx context.Context, req domain.UserAuthByGoogleRequest) (domain.UserAuthByGoogleResponse, error)
	UserAuthByLine(ctx context.Context, req domain.UserAuthByLineRequest) (domain.UserAuthByLineResponse, error)
}
