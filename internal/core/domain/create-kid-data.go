package domain

type CreateKidDataRequest struct {
	AccessToken          string
	PregnantId           string `json:"pregnant_id"`
	KidName              string `json:"kid_name"`
	KidBirthday          string `json:"kid_birthday"`
	KidGender            string `json:"kid_gender"`
	KidGestationalAge    string `json:"kid_gestational_age"`
	KidOxygen            string `json:"kid_oxygen"`
	KidCongenitalDisease string `json:"kid_congenital_disease"`
	KidBirthWeight       string `json:"kid_birth_weight"`
	KidBodyLength        string `json:"kid_body_length"`
	KidBloodType         string `json:"kid_blood_type"`
	UserId               string
}

type CreateKidDataResponse struct {
	Code    int    `json:"code" example:"200"`
	Message string `json:"message" example:"success"`
}
