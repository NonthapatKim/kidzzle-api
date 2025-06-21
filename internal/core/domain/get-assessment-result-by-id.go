package domain

import (
	"time"
)

type GetAssessmentResultByIdRequest struct {
	AccessToken          string
	KidId                string
	AssessmentQuestionId string
}

type GetAssessmentResultByIdResponse struct {
	AssessmentResultId   string    `json:"assessment_result_id" example:"8457529e-01fb-41d0-82b1-9618be667862"`
	KidId                string    `json:"kid_id" example:"7d7bc87f-0faf-43f9-9af9-1a09166db993"`
	AssessmentQuestionId string    `json:"assessment_question_id" example:"3afe2697-2b02-46d4-a9c3-efd0bdd9e014"`
	IsPassed             bool      `json:"is_passed" example:"true"`
	CreatedAt            time.Time `json:"created_at" example:"2025-03-09T00:00:00Z"`
}
