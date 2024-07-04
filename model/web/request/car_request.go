package request

type CarRequest struct {
	Brand string `json:"brand" binding:"required"`
	Model string `json:"model" binding:"required"`
	Year  int    `json:"year" binding:"required,min=1878"`
	Image string `json:"image" binding:"required,url"`
}
