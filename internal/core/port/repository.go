package port

import "github.com/NonthapatKim/kidzzle-api/internal/core/domain"

type Repository interface {
	// User
	CheckUser(req domain.CheckUserRequest) (domain.CheckUserResult, error)
	UserLogin(req domain.UserLoginRequest) (domain.UserLoginResult, error)
	UserLoginBySocial(req domain.UserLoginBySocialRequest) (domain.UserLoginBySocialResult, error)
	UserRegister(req domain.UserRegisterRequest) (domain.UserRegisterResponse, error)
	UserRegisterBySocial(req domain.UserRegisterBySocialRequest) (domain.UserRegisterBySocialResponse, error)
	UserRequestResetPassword(req domain.UserRequestResetPasswordRequest) (domain.UserRequestResetPasswordResult, error)
	UserResetPassword(req domain.UserResetPasswordRequest) (domain.UserResetPasswordResult, error)

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
}
