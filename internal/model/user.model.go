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

type UpdateRoleInput struct {
	UserID int    `json:"user_id" binding:"required,numeric,min=0"`
	Role   string `json:"role" binding:"required,role"`
}

type GetUsersForAdminInput struct {
	Limit  int    `json:"limit" binding:"required,numeric,max=20,gt=0"`
	Page   int    `json:"page" binding:"required,numeric,gt=0"`
	Role   string `json:"role" binding:"role"`
	Search string `json:"search" binding:"email_prefix,max=255,endsnotwith= ,startsnotwith= "`
}
type GetUsersForAdminOutput struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
}

type DeleteUserInput struct {
	UserID int `json:"user_id" binding:"required,numeric,min=0"`
}
