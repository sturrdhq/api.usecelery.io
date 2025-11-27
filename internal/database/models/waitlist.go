package models

import "github.com/google/uuid"

type WaitList struct {
	BaseModel
	Name          string         `gorm:"unique"`
	Subscriptions []Subscription `gorm:"foreignKey:WaitListID"`
}

type Subscription struct {
	BaseModel
	WaitListID uuid.UUID `json:"waitListId"`
	Email      string    `gorm:"unique"`
}
