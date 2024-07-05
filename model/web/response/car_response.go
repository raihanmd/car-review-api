package response

type CarResponse struct {
	ID    uint   `json:"id" example:"1"`
	Brand string `json:"brand" example:"Toyota"`
	Model string `json:"model" example:"Yaris"`
	Year  int    `json:"year" example:"2020"`
	Image string `json:"image" example:"url"`
}
