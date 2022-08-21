package repository

import (
	"MBETakeHomeTest/entity"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"log"
)

type UserBalanceRepository interface {
	InsertUserBalance(balance entity.UserBalance) entity.UserBalance
	UpdateUserBalance(balance entity.UserBalance) entity.UserBalance
	UserCurrentBalance(userid string) entity.UserBalance
	IsEnoughBalance(userid string,required_balance float64) (tx *gorm.DB)
}

type userBalanceConnection struct {
	connection *gorm.DB
}


func NewUserBalanceRepository(db *gorm.DB) UserBalanceRepository {
	return &userBalanceConnection{
		connection: db,
	}
}

func (db userBalanceConnection) IsEnoughBalance(userid string,required_balance float64) (tx *gorm.DB) {
	var balance entity.UserBalance
	return db.connection.Where("user_id = ? AND balance >= ?", userid,required_balance).Take(&balance)
}


func (db userBalanceConnection) InsertUserBalance(balance entity.UserBalance) entity.UserBalance {
	log.Println("balance to insert")
	guid,err := uuid.NewV4()
	if err != nil {
		panic("Error generate ID")
	}

	balance.ID = guid
	res := db.connection.Save(balance)
	if res.Error != nil{
		panic(res.Error.Error())
	}
	return balance
}

func (db userBalanceConnection) UpdateUserBalance(balance entity.UserBalance) entity.UserBalance {
	db.connection.Save(&balance)
	db.connection.Preload("User").Find(&balance)
	return balance
}

func (db userBalanceConnection) UserCurrentBalance(userid string) entity.UserBalance {
	var balance entity.UserBalance
	db.connection.Where("user_id = ?", userid).Take(&balance)
	return balance
}