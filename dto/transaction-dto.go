package dto

import (
	"github.com/gofrs/uuid"
	"time"
)

type TransactionDTO struct {
	UserIDFrom      uuid.UUID `json:"user_id_from" binding:"required"`
	UserIDTo        uuid.UUID `json:"user_id_to"`
	TransactionType int       `json:"transaction_type"`
	Flag            string    `json:"flag"`
	CreatedAt       time.Time `json:"transaction_time"`
	Balance         float64   `json:"balance" binding:"required"`
}
