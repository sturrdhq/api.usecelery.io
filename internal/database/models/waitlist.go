package models

import "github.com/google/uuid"

type WaitList struct {
	BaseModel
	Subscription []Subscription `gorm:"foreignKey:WaitlistID"`
	Name         string         `gorm:"unique"`
}

type Subscription struct {
	BaseModel
	WaitListID uuid.UUID `json:"waitListId" redis:"waitListId"`
	Email      string    `gorm:"email,unique"`
}
