package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Menu struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       string             `json:"name" validate:"required,min=2,max=100"`
	Category   string             `json:"category" validate:"required"`
	Start_date string             `json:"start_date"`
	End_date   string             `json:"end_date"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"Updated_at"`
	Menu_id    string             `json:"menu_id"`
}
