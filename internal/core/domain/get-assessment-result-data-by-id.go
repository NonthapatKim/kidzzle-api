package domain

import "time"

type GetAssessmentResultDataByIdRequest struct {
	AccessToken      string
	KidId            string
	AgeRageId        string
	AssessmentTypeId string
}

type GetAssessmentResultDataByIdResponse struct {
	AssessmentResultId   string    `json:"assessment_result_id" example:"7d7bc87f-0faf-43f9-9af9-1a09166db993"`
	KidId                string    `json:"kid_id" example:"8457529e-01fb-41d0-82b1-9618be667862"`
	AssessmentQuestionId string    `json:"assessment_question_id" example:"3afe2697-2b02-46d4-a9c3-efd0bdd9e014"`
	AssessmentTypeId     string    `json:"assessment_type_id" example:"ASSMTT_1"`
	AgeRangeId           string    `json:"age_range_id" example:"AGER_1"`
	IsPassed             bool      `json:"is_passed" example:"true"`
	CreatedAt            time.Time `json:"created_at" example:"2025-03-09T00:00:00Z"`
}
