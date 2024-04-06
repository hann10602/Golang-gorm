package common

import "time"

type SQLModel struct {
	Id        int        `json:"id"`
	CreatedAt *time.Time `json:"createdAt,omitempty" bson:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty" bson:"updatedAt"`
}
