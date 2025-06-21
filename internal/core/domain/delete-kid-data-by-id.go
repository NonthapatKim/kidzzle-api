package domain

type DeleteKidDataByIdRequest struct {
	AccessToken string
	KidId       string
}

type DeleteKidDataByIdResponse struct {
	Code    int    `json:"code" example:"200"`
	Message string `json:"message" example:"success"`
}
