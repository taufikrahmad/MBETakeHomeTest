package entity

import (
	"github.com/gofrs/uuid"
)

type UserBalance struct {
	ID      uuid.UUID `gorm:"type:uuid;primary_key:auto_increment" json:"id"`
	UserID  uuid.UUID `gorm:"type:uuid" json:"user_id"`
	Balance float64   `gorm:"type:decimal(22,2);"`
	User    User      `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
}
