package request

type ReviewCreateRequest struct {
	CarID   uint   `json:"car_id" binding:"required" extensions:"x-order=0"`
	Title   string `json:"title" binding:"required,max=100" extensions:"x-order=1"`
	Content string `json:"content" binding:"required" extensions:"x-order=2"`
	Image   string `json:"image" binding:"required,url" extensions:"x-order=3"`
}

type ReviewUpdateRequest struct {
	Title   *string `json:"title" binding:"omitempty" extensions:"x-order=0"`
	Content *string `json:"content" binding:"omitempty" extensions:"x-order=1"`
	Image   *string `json:"image" binding:"required,url" extensions:"x-order=2"`
}
