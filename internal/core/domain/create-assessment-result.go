package domain

type CreateAssessmentResultRequest struct {
	AccessToken          string
	KidId                string `json:"kid_id"`
	AssessmentQuestionId string `json:"assessment_question_id"`
	IsPassed             bool   `json:"is_passed"`
}

type CreateAssessmentResultResponse struct {
	Code    int    `json:"code" example:"200"`
	Message string `json:"message" example:"success"`
}
