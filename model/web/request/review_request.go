package request

type ReviewCreateRequest struct {
	CarID   uint   `json:"car_id" binding:"required"`
	Title   string `json:"title" binding:"required" binding:"max=100"`
	Content string `json:"content" binding:"required"`
}

type ReviewUpdateRequest struct {
	Title   *string `json:"title" binding:"omitempty"`
	Content *string `json:"content" binding:"omitempty"`
}
