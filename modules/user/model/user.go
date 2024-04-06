package model

import (
	"errors"
	"gin_mysql/common"
	"time"
)

var (
	ErrUsernameIsBlank = errors.New("Username cannot be blank")
	ErrPasswordIsBlank = errors.New("Password cannot be blank")
)

type Users struct {
	common.SQLModel
	Username string `json:"username" bson:"username" gorm:"size:255;unique"`
	Password string `json:"password" bson:"password" gorm:"size:255;unique"`
	Role     string `json:"role" bson:"role"`
	Status   string `json:"status" bson:"status"`
}

type CreateUserDTO struct {
	Id        int        `json:"id" gorm:"column:id"`
	Username  string     `json:"username" gorm:"column:username"`
	Password  string     `json:"password" gorm:"column:password"`
	Role      string     `json:"role" gorm:"default:USER;not null;column:role"`
	Status    string     `json:"status" gorm:"column:status;default:ACTIVE"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"column:created_at"`
}

type UpdateUserDTO struct {
	Username  *string    `json:"username,omitempty" gorm:"column:username"`
	Password  *string    `json:"password,omitempty" gorm:"column:password"`
	Role      *string    `json:"role,omitempty" gorm:"default:USER;not null;column:role"`
	Status    *string    `json:"status,omitempty" gorm:"column:status"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`
}
