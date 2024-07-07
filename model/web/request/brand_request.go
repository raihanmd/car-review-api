package request

type BrandRequest struct {
	Name string `json:"name" binding:"required,max=100"`
}
