package request

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=20,no_space,lowercase"`
	Password string `json:"password" binding:"required,min=6,no_space"`
}
