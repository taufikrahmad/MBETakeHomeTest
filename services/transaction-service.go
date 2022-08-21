package services

import (
	"MBETakeHomeTest/dto"
	"MBETakeHomeTest/entity"
	"MBETakeHomeTest/repository"
	"github.com/mashingan/smapping"
	"log"
	"time"
)

type TransactionService interface {
	CreateTransaction(transaction dto.TransactionDTO) entity.Transaction
	UserTransactions(userID string) []entity.TransactionHistory
	IsEnoughBalance(userID string, required_balance float64) bool
	IsDestinationExist(userID string) bool
}

type transactionService struct {
	transactionRepository repository.TransactionRepository
	balanceRepository     repository.UserBalanceRepository
	userRepository        repository.UserRepository
}

func (t transactionService) IsDestinationExist(userID string) bool {
	res := t.userRepository.IsUserExists(userID)
	return !(res.Error == nil)
}

func (t transactionService) CreateTransaction(transaction dto.TransactionDTO) entity.Transaction {
	transactionToInsert := entity.Transaction{}
	log.Println(transactionToInsert)
	err := smapping.FillStruct(&transactionToInsert, smapping.MapFields(&transaction))

	switch transaction.TransactionType {
	case 1:
		/*
			top up
			set user to = user from
		*/
		transaction.UserIDTo = transaction.UserIDFrom
		transaction.Flag = "D"
	case 2:
		/*
			with draw
			set user to = user from
		*/
		transaction.UserIDTo = transaction.UserIDFrom
		transaction.Flag = "C"
	case 3:
		/*
			transfer
			set user to = user from
		*/
		transaction.Flag = "C"
	default:

	}

	if err != nil {
		log.Fatal("Failed map %v", err)
	}

	transactionToInsert.CreatedAt = time.Now()
	insertTransaction := t.transactionRepository.InsertTransaction(transactionToInsert)

	//

	return insertTransaction
}

func (t transactionService) UserTransactions(userID string) []entity.TransactionHistory {
	return t.transactionRepository.UserTransaction(userID)
}

func (t transactionService) IsEnoughBalance(userID string, required_balance float64) bool {
	res := t.balanceRepository.IsEnoughBalance(userID, required_balance)
	return !(res.Error == nil)
}

func NewTransactionService(transactionRepository repository.TransactionRepository, balanceRepository repository.UserBalanceRepository,userRepository repository.UserRepository) TransactionService {
	return &transactionService{
		transactionRepository: transactionRepository,
		balanceRepository:     balanceRepository,
		userRepository : userRepository,
	}
}
