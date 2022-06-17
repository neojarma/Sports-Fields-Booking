package response

type UserResponse struct {
	IdUser       string `json:"idUser"`
	Name         string `json:"name"`
	Address      string `json:"address"`
	PhoneNumber  string `json:"phoneNumber"`
	Email        string `json:"email"`
	ImageProfile string `json:"imageProfile"`
}
