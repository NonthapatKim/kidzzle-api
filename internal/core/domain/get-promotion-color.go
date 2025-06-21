package domain

type GetPromotionColorResponse struct {
	ColorId           string `json:"color_id" example:"bd5f5f6f-fd05-11ef-870e-0242ac120002"`
	ColorName         string `json:"color_name" example:"สีน้ำตาล (Brown)"`
	ColorShape        string `json:"color_shape" example:"สี่เหลี่ยมจัตุรัส (Square)"`
	ColorImageUrl     string `json:"color_image_url" example:"url_link"`
	ColorSoundFileUrl string `json:"color_sound_file_url" example:"url_link"`
}
