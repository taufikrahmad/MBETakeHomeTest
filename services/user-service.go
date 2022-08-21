package services

import (
	"MBETakeHomeTest/dto"
	"MBETakeHomeTest/entity"
	"MBETakeHomeTest/repository"
	"github.com/mashingan/smapping"
	"log"
)

type UserService interface {
	Update(user dto.UserUpdateDTO) entity.User
	Profile(userID string) entity.User
}

type userService struct {
	userRepository repository.UserRepository
}

func (service userService) Update(user dto.UserUpdateDTO) entity.User {
	userToUpdate := entity.User{}
	log.Println(user)
	err := smapping.FillStruct(&userToUpdate,smapping.MapFields(&user))
	if err != nil{
		log.Fatal("Failed map %v",err)
	}

	updatedUser := service.userRepository.UpdateUser(userToUpdate)
	return updatedUser
}

func (service userService) Profile(userID string) entity.User {
	return service.userRepository.ProfileUser(userID)
}

func NewUserService(userRepo repository.UserRepository) UserService{
	return &userService{
		userRepository: userRepo,
	}
}
