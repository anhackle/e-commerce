package model

type GetProfileOutput struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
}

type UpdateProfileInput struct {
	FirstName   string `json:"first_name" binding:"name,max=50"`
	LastName    string `json:"last_name" binding:"name,max=50"`
	PhoneNumber string `json:"phone_number" binding:"phone,max=15"`
	Address     string `json:"address" binding:"name,max=255"`
}

type ChangePasswordInput struct {
	OldPassword     string `json:"old_password" binding:"password"`
	NewPassword     string `json:"new_password" binding:"password"`
	ConfirmPassword string `json:"confirm_password" binding:"password"`
}
