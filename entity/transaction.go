package entity

import (
	"github.com/gofrs/uuid"
	"time"
)

type Transaction struct {
	ID              uuid.UUID `gorm:"type:uuid;primary_key:auto_increment" json:"id"`
	UserIDFrom      uuid.UUID `gorm:"type:uuid" json:"user_id_from"`
	UserIDTo        uuid.UUID `gorm:"type:uuid" json:"user_id_to"`
	TransactionType int       `gorm:"type:int" json:"transaction_type"`
	Flag            string    `gorm:"type:char" json:"flag"`
	CreatedAt       time.Time `json:"transaction_time"`
	Balance         float64   `gorm:"type:decimal(22,2);"`
	UserFrom        User      `gorm:"foreignkey:UserIDFrom;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
	UserTo          User      `gorm:"foreignkey:UserIDTo;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user_to"`
}

type TransactionHistory struct {
	ID              uuid.UUID `gorm:"type:uuid;primary_key:auto_increment" json:"id"`
	UserIDFrom      uuid.UUID `gorm:"type:uuid" json:"user_id_from"`
	UserIDTo        uuid.UUID `gorm:"type:uuid" json:"user_id_to"`
	TransactionType int       `gorm:"type:int" json:"transaction_type"`
	TransactionDesc string       `gorm:"type:char" json:"transaction_desc"`
	Flag            string    `gorm:"type:char" json:"flag"`
	CreatedAt       time.Time `json:"transaction_time"`
	Balance         float64   `gorm:"type:decimal(22,2);"`
	UserFrom        string    `json:"user_from"`
	UserTo          string    `json:"user_to"`
}
