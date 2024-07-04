package request

type UpdateUserProfileRequest struct {
	Username string  `json:"username" binding:"omitempty,min=3,max=20,no_space,lowercase"`
	Email    *string `json:"email" binding:"omitempty,email"`
	FullName *string `json:"full_name" binding:"omitempty,min=3,max=100"`
	Bio      *string `json:"bio" binding:"omitempty,max=500"`
	Age      *int    `json:"age" binding:"omitempty,min=0"`
	Gender   *string `json:"gender" binding:"omitempty,uppercase,oneof=MALE FEMALE"`
}
