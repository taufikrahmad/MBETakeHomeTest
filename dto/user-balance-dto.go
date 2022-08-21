package dto

import (
	"github.com/gofrs/uuid"
)

type UserBalanceDTO struct {
	UserID uuid.UUID `json:"user_id" form:"user_id"`
	Balance float64 `gorm:"type:decimal(22,2);"`
}

type UserBalanceUpdateDTO struct {
	ID uuid.UUID `json:"id" binding:"required" form:"id"`
	UserID uuid.UUID `json:"user_id" form:"user_id"`
	Balance float64 `gorm:"type:decimal(22,2);"`
}
