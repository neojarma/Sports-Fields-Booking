package request

type RegisterRequest struct {
	Name        string `json:"name" validate:"required"`
	Address     string `json:"address" validate:"required"`
	PhoneNumber string `json:"phoneNumber" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Username    string `json:"username"`
	Password    string `json:"password"`
}
