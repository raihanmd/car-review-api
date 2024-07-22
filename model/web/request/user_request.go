package request

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=20,no_space,lowercase" extensions:"x-order=0"`
	Email    string `json:"email" binding:"required,email" extensions:"x-order=1"`
	Password string `json:"password" binding:"required,min=8,no_space" example:"password" extensions:"x-order=2"`
}

type ForgotPasswordRequest struct {
	Username string `json:"username" binding:"required,min=3,max=20,no_space,lowercase" extensions:"x-order=0"`
	Email    string `json:"email" binding:"required,email"`
}

type ResetPasswordRequest struct {
	Token       string `json:"token" binding:"required" example:"token" extensions:"x-order=0"`
	NewPassword string `json:"new_password" binding:"required,min=8" example:"new_password" extensions:"x-order=1"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email" extensions:"x-order=0"`
	Password string `json:"password" binding:"required" example:"password" extensions:"x-order=1"`
}

type UpdatePasswordRequest struct {
	Password string `json:"password" binding:"required,min=3,no_space"`
}

type UpdateUserProfileRequest struct {
	Username string  `json:"username" binding:"omitempty,min=3,max=20,no_space,lowercase" extensions:"x-order=0"`
	Email    string  `json:"email" binding:"omitempty,email" extensions:"x-order=1"`
	FullName *string `json:"full_name" binding:"omitempty,min=3,max=100" extensions:"x-order=2"`
	Bio      *string `json:"bio" binding:"omitempty,max=500" extensions:"x-order=3"`
	Age      *int    `json:"age" binding:"omitempty,min=0" extensions:"x-order=4"`
	Gender   *string `json:"gender" binding:"omitempty,uppercase,oneof=MALE FEMALE" extensions:"x-order=5"`
}
