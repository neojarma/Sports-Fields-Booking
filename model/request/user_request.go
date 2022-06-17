package request

type UserRequest struct {
	IdUser       string `json:"idUser,omitempty"`
	Name         string `json:"name"`
	Address      string `json:"address"`
	PhoneNumber  string `json:"phoneNumber"`
	Email        string `json:"email"`
	ImageProfile string `json:"profile,omitempty"`
	Username     string `json:"username,omitempty"`
	Password     string `json:"password,omitempty"`
}
