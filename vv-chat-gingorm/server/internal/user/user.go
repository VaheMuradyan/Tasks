package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Email    string `gorm:"unique"`
	Password string
}

type CreateUserReq struct {
	Username string
	Email    string `gorm:"unique"`
	Password string
}

type LoginUserReq struct {
	Email    string `gorm:"unique"`
	Password string
}
