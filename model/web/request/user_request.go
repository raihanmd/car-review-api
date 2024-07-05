package request

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=20,no_space,lowercase" extensions:"x-order=0"`
	Password string `json:"password" binding:"required,min=6,no_space" extensions:"x-order=1"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required" extensions:"x-order=0"`
	Password string `json:"password" binding:"required" extensions:"x-order=1"`
}

type UpdatePasswordRequest struct {
	Password string `json:"password" binding:"required,min=3,no_space"`
}

type UpdateUserProfileRequest struct {
	Username string  `json:"username" binding:"omitempty,min=3,max=20,no_space,lowercase" extensions:"x-order=0"`
	Email    *string `json:"email" binding:"omitempty,email" extensions:"x-order=1"`
	FullName *string `json:"full_name" binding:"omitempty,min=3,max=100" extensions:"x-order=2"`
	Bio      *string `json:"bio" binding:"omitempty,max=500" extensions:"x-order=3"`
	Age      *int    `json:"age" binding:"omitempty,min=0" extensions:"x-order=4"`
	Gender   *string `json:"gender" binding:"omitempty,uppercase,oneof=MALE FEMALE" extensions:"x-order=5"`
}
