package domain

type UserAuthenticateRequest struct {
	AccessToken string
}

type UserAuthenticateResponse struct {
	UserId string `json:"user_id" example:"302ba3d3-7c92-4777-a914-2aa7dee55e1e"`
}
