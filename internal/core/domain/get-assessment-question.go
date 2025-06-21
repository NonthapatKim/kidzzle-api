package domain

type GetAssessmentQuestionRequest struct {
	AssessmentTypeId string
}

type GetAssessmentQuestionResponse struct {
	AssessmentQuestionId    string  `json:"assessment_question_id" example:"088e7d6b-faa8-11ef-870e-0242ac120002"`
	AssessmentTypeId        string  `json:"assessment_type_id" example:"ASSMTT_1"`
	AgeRangeId              string  `json:"age_range_id" example:"AGER_4"`
	DevTypeId               string  `json:"dev_type_id" example:"DEVT_1"`
	AssessmentNo            string  `json:"assessment_no" example:"1"`
	AssessmentTypeName      string  `json:"assessment_type_name" example:"DAIM"`
	AgeRangeName            string  `json:"age_range_name" example:"3-4"`
	DevTypeName             string  `json:"development_type" example:"GM"`
	QuestionText            string  `json:"question_text" example:"ท่านอนคว่ำยกศีรษะและอกพ้นพื้น (GM)"`
	AssessmentMethod        string  `json:"assessment_method" example:"<p>1. จัดให้เด็กอยู่ในท่านอนคว่ำบนพื้นราบ&nbsp;</p><p>2. เขย่ากรุ๋งกริ๋งด้านหน้าเด็กเพื่อให้เด็กสนใจ แล้วเคลื่อนขึ้นด้านบน กระตุ้นให้เด็กมองตาม</p>"`
	AssessmentRequiredtTool *string `json:"assessment_required_tool" example:"กรุ๋งกริ๋ง"`
	PassCriteria            string  `json:"pass_criteria" example:"เด็กยกศีรษะและอกโดยใช้แขนยันกับพื้นพยุงตัวไว้อย่างน้อย 5 วินาที"`
}
