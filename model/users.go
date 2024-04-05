package model

import "time"

type Users struct {
	Id        int        `json:"id" gorm:"column:id;primary_key"`
	Username  string     `json:"username" bson:"username" gorm:"size:255;unique"`
	Password  string     `json:"password" bson:"password" gorm:"size:255;unique"`
	Role      string     `json:"role" bson:"role"`
	Status    string     `json:"status" bson:"status"`
	CreatedAt *time.Time `json:"createdAt,omitempty" bson:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty" bson:"updatedAt"`
}
