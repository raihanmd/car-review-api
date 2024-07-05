package response

type ReviewResponse struct {
	ID      uint   `json:"id" example:"1"`
	CarID   uint   `json:"car_id" example:"2"`
	UserID  uint   `json:"user_id" example:"3"`
	Title   string `json:"title" example:"Title"`
	Content string `json:"content" example:"Lorem ipsum dolor sit amet"`
}

type FindReviewResponse struct {
	ID      uint               `json:"id" example:"1"`
	Car     CarResponse        `json:"car"`
	User    ReviewUserResponse `json:"user"`
	Title   string             `json:"title" example:"Title"`
	Content string             `json:"content" example:"Lorem ipsum dolor sit amet"`
}

type ReviewUserResponse struct {
	ID       uint   `json:"id" example:"1"`
	Username string `json:"username" example:"John Doe"`
}
