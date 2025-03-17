package model

type RegisterInput struct {
	Email    string `json:"email" binding:"required,email,lowercase,max=50"`
	Password string `json:"password" binding:"required,password,max=50"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email,lowercase,max=50"`
	Password string `json:"password" binding:"required,password,max=50"`
}
