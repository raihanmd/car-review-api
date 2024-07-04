package request

type UpdatePasswordRequest struct {
	Password string `json:"password" binding:"required,min=3,no_space"`
}
