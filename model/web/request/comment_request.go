package request

type CommentCreateRequest struct {
	ReviewID uint   `json:"review_id" binding:"required" extensions:"x-order=0"`
	Content  string `json:"content" binding:"required" extensions:"x-order=1"`
}

type CommentUpdateRequest struct {
	Content string `json:"content" binding:"required" extensions:"x-order=0"`
}
