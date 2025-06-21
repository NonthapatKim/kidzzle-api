package domain

type UpdateMotherPregnantByIdRequest struct {
	AccessToken               string
	PregnantId                string
	MotherName                string `json:"mother_name"`
	MotherBirthday            string `json:"mother_birthday"`
	PregnantCongenitalDisease string `json:"pregnant_congenital_disease"`
	PregnantComplications     string `json:"pregnant_complications"`
	PregnantDrugHistory       string `json:"pregnant_drug_history"`
}

type UpdateMotherPregnantByIdResponse struct {
	Code    int    `json:"code" example:"200"`
	Message string `json:"message" example:"success"`
}
