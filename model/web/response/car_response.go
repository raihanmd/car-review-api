package response

type CarResponse struct {
	ID    uint   `json:"id" example:"1" extensions:"x-order=0"`
	Brand string `json:"brand" example:"Toyota" extensions:"x-order=1"`
	Model string `json:"model" example:"Yaris" extensions:"x-order=2"`
	Year  int    `json:"year" example:"2020" extensions:"x-order=3"`
	Image string `json:"image" example:"url" extensions:"x-order=4"`
}
