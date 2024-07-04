package response

type RegisterResponse struct {
	Username string `json:"username" example:"luigi"`
	Role     string `json:"role" example:"USER"`
}
