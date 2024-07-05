package request

type CarCreateRequest struct {
	Brand string `json:"brand" binding:"required" extensions:"x-order=0"`
	Model string `json:"model" binding:"required" extensions:"x-order=1"`
	Year  int    `json:"year" binding:"required,min=1878" extensions:"x-order=2"`
	Image string `json:"image" binding:"required,url" extensions:"x-order=3"`
}

type CarUpdateRequest struct {
	Brand *string `json:"brand" binding:"omitempty" extensions:"x-order=0"`
	Model *string `json:"model" binding:"omitempty" extensions:"x-order=1"`
	Year  *int    `json:"year" binding:"omitempty,min=1878" extensions:"x-order=2"`
	Image *string `json:"image" binding:"omitempty,url" extensions:"x-order=3"`
}
