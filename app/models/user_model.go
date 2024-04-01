package models

type SignUpRequest struct {
	Email    string `json:"email" validate:"required,email,lte=255"`
	Password string `json:"password" validate:"required,lte=255"`
	UserRole string `json:"userRole" validate:"required,lte=25"`
}

type SignUpResponse struct {
	ID                int    `json:"id"`
	Email             string `json:"email"`
	Status            int    `json:"status"`
	UserRole          string `json:"userRole"`
	OldID             int    `json:"oldId"`
	MobilePhoneNumber string `json:"mobilePhoneNumber"`
}
