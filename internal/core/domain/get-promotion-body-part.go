package domain

type GetPromotionBodyPartResponse struct {
	BodyId           string `json:"body_id" example:"bd5f5f6f-fd05-11ef-870e-0242ac120002"`
	BodyName         string `json:"body_name" example:"สีเขียว (Green)"`
	BodyImageUrl     string `json:"body_image_url" example:"url_link"`
	BodySoundFileUrl string `json:"body_sound_file_url" example:"url_link"`
}
