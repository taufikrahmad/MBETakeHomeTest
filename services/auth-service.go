package services

import (
	"MBETakeHomeTest/dto"
	"MBETakeHomeTest/entity"
	"MBETakeHomeTest/repository"
	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type AuthService interface {
	VerifyCredential(email string,password string) interface{}
	CreateUser(user dto.RegisterDTO) entity.User
	FindUserByEmail(email string) entity.User
	IsDuplicateEmail(email string) bool
}

type authService struct {
	userRepository repository.UserRepository
}

func (service authService) IsDuplicateEmail(email string) bool {
	res := service.userRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

func (service authService) VerifyCredential(email string, password string) interface{} {
	res := service.userRepository.VerifyCredential(email,password)
	if v, ok := res.(entity.User); ok{
		comparedPassword := comparePassword(v.Password, []byte(password))
		if v.Email == email && comparedPassword {
			return res
		}
		return false
	}

	return false
}

func (service authService) CreateUser(user dto.RegisterDTO) entity.User {
	userToCreate := entity.User{}
	log.Println(user)
	err := smapping.FillStruct(&userToCreate,smapping.MapFields(&user))
	if err != nil {
		log.Fatal("Failed mapping %v",err)
	}

	res := service.userRepository.InsertUser(userToCreate)
	return res

}

func (service authService) FindUserByEmail(email string) entity.User {
	return service.userRepository.FindUserByEmail(email)
}

func NewAuthService(userRep repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRep,
	}
}

func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}