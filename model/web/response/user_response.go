package response

type RegisterResponse struct {
	Username string `json:"username" example:"luigi" extensions:"x-order=0"`
	Role     string `json:"role" example:"USER" extensions:"x-order=1"`
}

type LoginResponse struct {
	Username string `json:"username" example:"luigi" extensions:"x-order=0"`
	Role     string `json:"role" example:"USER" extensions:"x-order=1"`
	Token    string `json:"token" example:"token" extensions:"x-order=2"`
}

type UserProfileResponse struct {
	ID       uint    `json:"id" example:"1" extensions:"x-order=0"`
	Username string  `json:"username" example:"luigi" extensions:"x-order=1"`
	Role     string  `json:"role" example:"USER" extensions:"x-order=2"`
	Email    *string `json:"email" example:"luigi@sam.com" extensions:"x-order=3"`
	FullName *string `json:"full_name" example:"Luigi Di Caprio" extensions:"x-order=4"`
	Bio      *string `json:"bio" example:"I am Luigi" extensions:"x-order=5"`
	Age      *int    `json:"age" example:"18" extensions:"x-order=6"`
	Gender   *string `json:"gender" example:"MALE" extensions:"x-order=7"`
}

type UpdateUserProfileResponse struct {
	ID       uint    `json:"id" example:"1" extensions:"x-order=0"`
	Username string  `json:"username" example:"luigi" extensions:"x-order=1"`
	Role     string  `json:"role" example:"USER" extensions:"x-order=2"`
	Email    *string `json:"email" example:"luigi@sam.com" extensions:"x-order=3"`
	FullName *string `json:"full_name" example:"Luigi Di Caprio" extensions:"x-order=4"`
	Bio      *string `json:"bio" example:"I am Luigi" extensions:"x-order=5"`
	Age      *int    `json:"age" example:"18" extensions:"x-order=6"`
	Gender   *string `json:"gender" example:"MALE" extensions:"x-order=7"`
}
