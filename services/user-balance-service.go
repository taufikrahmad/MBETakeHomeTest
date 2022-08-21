package services

import (
	"MBETakeHomeTest/dto"
	"MBETakeHomeTest/entity"
	"MBETakeHomeTest/repository"
	"github.com/mashingan/smapping"
	"log"
)

type UserBalanceService interface {
	InsertUserBalance(balance dto.UserBalanceDTO) entity.UserBalance
	UpdateUserBalance(balance dto.UserBalanceUpdateDTO) entity.UserBalance
	UserCurrentBalance(userid string) entity.UserBalance
}

type userBalanceService struct {
	userBalanceRepository repository.UserBalanceRepository
}

func NewUserBalanceService(userRep repository.UserBalanceRepository) UserBalanceService {
	return &userBalanceService{
		userBalanceRepository: userRep,
	}
}

func (service userBalanceService) InsertUserBalance(balance dto.UserBalanceDTO) entity.UserBalance {
	userBalanceToUpdate := entity.UserBalance{}
	log.Println(balance)
	err := smapping.FillStruct(&userBalanceToUpdate,smapping.MapFields(&balance))
	if err != nil{
		log.Fatal("Failed map %v",err)
	}

	updatedUser := service.userBalanceRepository.InsertUserBalance(userBalanceToUpdate)
	return updatedUser
}

func (service userBalanceService) UpdateUserBalance(balance dto.UserBalanceUpdateDTO) entity.UserBalance {
	userBalanceToUpdate := entity.UserBalance{}
	log.Println(balance)
	err := smapping.FillStruct(&userBalanceToUpdate,smapping.MapFields(&balance))
	if err != nil{
		log.Fatal("Failed map %v",err)
	}

	updatedUser := service.userBalanceRepository.UpdateUserBalance(userBalanceToUpdate)
	return updatedUser
}

func (service userBalanceService) UserCurrentBalance(userid string) entity.UserBalance {
	return service.userBalanceRepository.UserCurrentBalance(userid)
}
