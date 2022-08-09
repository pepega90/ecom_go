package users

import (
	"ecom_go/models/carts"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string       `json:"name"`
	Email    string       `json:"email"`
	Phone    int          `json:"phone"`
	Password []byte       `json:"password"`
	Carts    []carts.Cart `json:"carts"`
}

func (u *User) HashPassword(pass string) {
	hashedPass, _ := bcrypt.GenerateFromPassword([]byte(pass), 14)
	u.Password = hashedPass
}

func (u *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(u.Password, []byte(password))
}
