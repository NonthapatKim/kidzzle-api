package domain

type DeleteMotherPregnantByIdRequest struct {
	AccessToken string
	PregnantId  string `json:"pregnant_id"`
}

type DeleteMotherPregnantByIdResponse struct {
	Code    int    `json:"code" example:"200"`
	Message string `json:"message" example:"success"`
}
