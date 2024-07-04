package request

type CarCreateRequest struct {
	Brand string `json:"brand" binding:"required"`
	Model string `json:"model" binding:"required"`
	Year  int    `json:"year" binding:"required,min=1878"`
	Image string `json:"image" binding:"required,url"`
}

type CarUpdateRequest struct {
	Brand *string `json:"brand" binding:"omitempty"`
	Model *string `json:"model" binding:"omitempty"`
	Year  *int    `json:"year" binding:"omitempty,min=1878"`
	Image *string `json:"image" binding:"omitempty,url"`
}
