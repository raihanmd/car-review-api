package response

type BrandResponse struct {
	ID   uint   `json:"id" example:"1" extensions:"x-order=0"`
	Name string `json:"name" exaple:"Toyota" extensions:"x-order=1"`
}
