package response

type FavouriteResponse struct {
	CarID    uint   `json:"car_id" example:"1" extensions:"x-order=0"`
	Brand    string `json:"brand" example:"Honda" extensions:"x-order=1"`
	Model    string `json:"model" example:"Civic" extensions:"x-order=2"`
	ImageUrl string `json:"image_url" example:"image url" extensions:"x-order=3"`
}
