package domain

import "time"

type GetMotherPregnantByIdRequest struct {
	AccessToken string
	UserId      string
}

type GetMotherPregnantByIdResponse struct {
	PregnantId                string     `json:"pregnant_id" example:"088e7d6b-faa8-11ef-870e-0242ac120002"`
	MotherName                string     `json:"mother_name" example:"TEST"`
	MotherBirthday            time.Time  `json:"mother_birthday" example:"2025-03-09T00:00:00Z"`
	PregnantCongenitalDisease string     `json:"pregnant_congenital_disease" example:"TEST"`
	PregnantComplications     string     `json:"pregnant_complications" example:"TEST"`
	PregnantDrugHistory       string     `json:"pregnant_drug_history" example:"TEST"`
	CreatedAt                 time.Time  `json:"created_at" example:"2025-03-09T00:00:00Z"`
	UpdatedAt                 *time.Time `json:"updated_at" example:"2025-03-09T00:00:00Z"`
}
