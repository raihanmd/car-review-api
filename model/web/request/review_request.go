package request

type ReviewQueryRequest struct {
	Title *string `form:"title" extensions:"x-order=0"`
	CarID *uint   `form:"car_id" extensions:"x-order=1"`
}

type ReviewCreateRequest struct {
	CarID    uint   `json:"car_id" binding:"required" extensions:"x-order=0"`
	Title    string `json:"title" binding:"required,max=100" extensions:"x-order=1"`
	Content  string `json:"content" binding:"required" extensions:"x-order=2"`
	ImageUrl string `json:"image_url" binding:"required,url" extensions:"x-order=3"`
}

type ReviewUpdateRequest struct {
	Title    *string `json:"title" extensions:"x-order=0"`
	Content  *string `json:"content" extensions:"x-order=1"`
	ImageUrl *string `json:"image_url" binding:"url" extensions:"x-order=2"`
}
