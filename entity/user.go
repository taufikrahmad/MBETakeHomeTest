package entity

import "github.com/gofrs/uuid"

type User struct {
	ID       uuid.UUID    `gorm:"type:uuid;primary_key:auto_increment" json:"id"`
	Name     string       `gorm:"type:varchar(255)" json:"name"`
	Email    string       `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password string       `gorm:"->;<-;not null" json:"-"`
	Token    string       `gorm:"-" json:"token,omitempty"`
	Balance  *UserBalance `json:"balance,omitempty"`
}
