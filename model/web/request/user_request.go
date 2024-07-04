package request

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=20,no_space,lowercase"`
	Password string `json:"password" binding:"required,min=6,no_space"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdatePasswordRequest struct {
	Password string `json:"password" binding:"required,min=3,no_space"`
}

type UpdateUserProfileRequest struct {
	Username string  `json:"username" binding:"omitempty,min=3,max=20,no_space,lowercase"`
	Email    *string `json:"email" binding:"omitempty,email"`
	FullName *string `json:"full_name" binding:"omitempty,min=3,max=100"`
	Bio      *string `json:"bio" binding:"omitempty,max=500"`
	Age      *int    `json:"age" binding:"omitempty,min=0"`
	Gender   *string `json:"gender" binding:"omitempty,uppercase,oneof=MALE FEMALE"`
}
