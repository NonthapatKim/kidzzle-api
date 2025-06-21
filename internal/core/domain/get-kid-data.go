package domain

import "time"

type GetKidDataRequest struct {
	AccessToken string
	PregnantId  string
	UserId      string
}

type GetKidDataResponse struct {
	KidId                string     `json:"kid_id" example:"088e7d6b-faa8-11ef-870e-0242ac120002"`
	PregnantId           string     `json:"pregnant_id" example:"088e7d6b-faa8-11ef-870e-0242ac120002"`
	MotherName           string     `json:"mother_name" example:"TEST"`
	KidName              string     `json:"kid_name" example:"Waruntorn Paonil"`
	KidBirthday          string     `json:"kid_birthday" example:"2025-03-09T00:00:00Z"`
	KidGender            string     `json:"kid_gender" example:"Male"`
	KidGestationalAge    string     `json:"kid_gestational_age" example:"TEST"`
	KidOxygen            string     `json:"kid_oxygen" example:"TEST"`
	KidBirthWeight       string     `json:"kid_birth_weight" example:"TEST"`
	KidBodyLength        string     `json:"kid_body_length" example:"TEST"`
	KidBloodType         string     `json:"kid_blood_type" example:"TEST"`
	KidCongenitalDisease string     `json:"kid_congenital_disease" example:"TEST"`
	CreatedAt            time.Time  `json:"created_at" example:"2025-03-09T00:00:00Z"`
	UpdatedAt            *time.Time `json:"updated_at" example:"2025-03-09T00:00:00Z"`
}
