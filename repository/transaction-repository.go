package repository

import (
	"MBETakeHomeTest/entity"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"time"
)

type TransactionRepository interface {
	InsertTransaction(transaction entity.Transaction) entity.Transaction
	UserTransaction(userid string) []entity.TransactionHistory
}

type transactionConnection struct {
	connection *gorm.DB
}


func NewUserTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionConnection{
		connection: db,
	}
}

func (db transactionConnection) InsertTransaction(transaction entity.Transaction) entity.Transaction {
	log.Println("balance to insert")
	guid,err := uuid.NewV4()
	if err != nil {
		panic("Error generate ID")
	}

	transaction.ID = guid
	transaction.CreatedAt = time.Now()
	res := db.connection.Save(transaction)
	if res.Error != nil{
		panic(res.Error.Error())
	}
	return transaction
}

func (db transactionConnection) UserTransaction(userid string) []entity.TransactionHistory {
	var transactions []entity.TransactionHistory
	db.connection.Debug().Table("transactions").Joins("left join users on users.id = transactions.user_id_from").Joins("left join users as user_to on user_to.id = transactions.user_id_from").Where("user_id_from = ?",userid).Select([]string{"transactions.*","CASE WHEN transactions.transaction_type = 1 THEN 'Top Up' WHEN transactions.transaction_type = 2 THEN 'With Draw' WHEN transactions.transaction_type = 3 THEN 'Transfer' ELSE '-' END AS transaction_desc", "user_to.name as user_to"}).Preload(clause.Associations).Scan(&transactions)
	return transactions
}