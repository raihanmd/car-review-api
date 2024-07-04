package response

type LoginResponse struct {
	Username string `json:"username" example:"luigi"`
	Role     string `json:"role" example:"USER"`
	Token    string `json:"token"`
}
