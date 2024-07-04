package response

type UserProfileResponse struct {
	ID       uint    `json:"id" example:"1"`
	Username string  `json:"username" example:"luigi"`
	Role     string  `json:"role" example:"USER"`
	Email    *string `json:"email" example:"luigi@sam.com"`
	FullName *string `json:"full_name" example:"Luigi Di Caprio"`
	Bio      *string `json:"bio" example:"I am Luigi"`
	Age      *int    `json:"age" example:"18"`
	Gender   *string `json:"gender" example:"MALE"`
}
