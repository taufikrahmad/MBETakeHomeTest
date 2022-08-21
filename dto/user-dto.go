package dto

import "github.com/gofrs/uuid"

type UserUpdateDTO struct {
	ID uuid.UUID `json:"id" form:"id"`
	Name string `json:"name" form:"name" binding:"required"`
	Email string `json:"email" form:"email" binding:"required" validate:"email"`
	Password string `json:"password,onitempty" form:"password,onitempty" validate:"min:6"`
}
