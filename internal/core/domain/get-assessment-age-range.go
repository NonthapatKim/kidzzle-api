package domain

type GetAssessmentAgeRangeRequest struct {
	AssessmentTypeId string `json:"assessment_type_id" example:"ASSMTT_1"`
}

type GetAssessmentAgeRangeResponse struct {
	AssessmentTypeId       *string `json:"assessment_type_id" example:"ASSMTT_1"`
	AgeRangeId             string  `json:"age_range_id" example:"AGER_1"`
	AgeRange               string  `json:"age_range" example:"0-1"`
	MinMonths              *string `json:"min_months" example:"0"`
	MaxMonths              *string `json:"max_months" example:"1"`
	AssessmentVideoUrl     *string `json:"assessment_video_url" example:"url"`
	AssessmentVideoUrlDspm *string `json:"-"`
	AssessmentVideoUrlDaim *string `json:"-"`
}
