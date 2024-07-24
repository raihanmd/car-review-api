package response

import "time"

type ReviewResponse struct {
	ID        uint      `json:"id" example:"1" extensions:"x-order=0"`
	CarID     uint      `json:"car_id" example:"2" extensions:"x-order=1"`
	UserID    uint      `json:"user_id" example:"3" extensions:"x-order=2"`
	Title     string    `json:"title" example:"Title" extensions:"x-order=3"`
	Content   string    `json:"content" example:"Lorem ipsum dolor sit amet" extensions:"x-order=4"`
	ImageUrl  string    `json:"image_url" example:"image url" extensions:"x-order=5"`
	CreatedAt time.Time `json:"created_at" example:"2022-01-01T00:00:00Z" extensions:"x-order=6"`
	UpdatedAt time.Time `json:"updated_at" example:"2022-01-01T00:00:00Z" extensions:"x-order=7"`
}

type FindReviewResponse struct {
	ID        uint               `json:"id" example:"1" extensions:"x-order=0"`
	Title     string             `json:"title" example:"Title" extensions:"x-order=1"`
	Content   string             `json:"content" example:"Lorem ipsum dolor sit amet" extensions:"x-order=2"`
	ImageUrl  string             `json:"image_url" example:"image url" extensions:"x-order=3"`
	User      ReviewUserResponse `json:"user" extensions:"x-order=5"`
	CreatedAt time.Time          `json:"created_at" example:"2022-01-01T00:00:00Z" extensions:"x-order=6"`
	UpdatedAt time.Time          `json:"updated_at" example:"2022-01-01T00:00:00Z" extensions:"x-order=7"`
}

type ReviewUserResponse struct {
	ID       uint   `json:"id" example:"1" extensions:"x-order=0"`
	Username string `json:"username" example:"John Doe" extensions:"x-order=1"`
}
