package users

type UserDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    int    `json:"phone"`
	Password string `json:"password"`
}
