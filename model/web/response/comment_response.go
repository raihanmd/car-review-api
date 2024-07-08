package response

import "time"

type CommentResponse struct {
	ID        uint                `json:"id" example:"1" extensions:"x-order=0"`
	ReviewID  uint                `json:"review_id" example:"2" extensions:"x-order=1"`
	User      CommentUserResponse `json:"user" extensions:"x-order=2"`
	Content   string              `json:"content" example:"Lorem ipsum dolor sit amet" extensions:"x-order=3"`
	CreatedAt time.Time           `json:"created_at" example:"2022-01-01T00:00:00Z"`
	UpdatedAt time.Time           `json:"updated_at" example:"2022-01-01T00:00:00Z"`
}

type CommentUserResponse struct {
	Username string `json:"username" example:"luigi" extensions:"x-order=0"`
}
