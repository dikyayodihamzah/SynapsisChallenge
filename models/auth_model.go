package models

type LoginRequest struct {
	Email    string `json:"email_req" binding:"required"`
	Password string `json:"pass_req" binding:"required"`
}

type SignUpRequest struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required, email" gorm:"unique"`
	Password        string `json:"password" binding:"required"`
	PasswordConfirm string `json:"password_confirm" binding:"required"`
}