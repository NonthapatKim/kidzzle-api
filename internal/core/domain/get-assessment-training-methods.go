package domain

type GetAssessmentTrainingMethodsRequest struct {
	AssessmentTypeId string
}

type GetAssessmentTrainingMethodsResponse struct {
	TrainingMethodsId     string  `json:"training_methods_id" example:"088e7d6b-faa8-11ef-870e-0242ac120002"`
	AssessmentQuestionId  string  `json:"assessment_question_id" example:"088e7d6b-faa8-11ef-870e-0242ac120002"`
	AssessmentTypeId      string  `json:"assessment_type_id" example:"ASSMTT_1"`
	AssessmentNo          string  `json:"assessment_no" example:"1"`
	TrainingText          string  `json:"training_text" example:"Training Text"`
	TrainingRequiredTools *string `json:"training_required_tools" example:"Training Required Tools"`
}
