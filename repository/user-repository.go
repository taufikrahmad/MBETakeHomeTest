package repository

import (
	"MBETakeHomeTest/entity"
	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type UserRepository interface {
	InsertUser(entity.User) entity.User
	UpdateUser(user entity.User) entity.User
	VerifyCredential(email string, password string) interface{}
	IsDuplicateEmail(email string) (tx *gorm.DB)
	IsUserExists(userid string) (tx *gorm.DB)
	FindUserByEmail(email string) entity.User
	ProfileUser(userid string) entity.User
}

type userConnection struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}


func (db userConnection) ProfileUser(userid string) entity.User {
	var user entity.User
	db.connection.Preload("Books").Preload("Books.User").First(&user,"id = ?", userid)
	return user
}

func (db userConnection) InsertUser(user entity.User) entity.User {
	guid,err := uuid.NewV4()
	if err != nil {
		panic("Error generate ID")
	}

	user.ID = guid
	user.Password = hashAndSalt([]byte(user.Password))
	log.Println("user to insert")
	log.Println(user)
	res := db.connection.Save(user)
	if res.Error != nil{
		panic(res.Error.Error())
	}
	return user
}

func (db userConnection) UpdateUser(user entity.User) entity.User {
	if user.Password != "" {
		user.Password = hashAndSalt([]byte(user.Password))
	} else {
		var tempUser entity.User
		db.connection.First(&user,"id = ?", user.ID)
		user.Password = tempUser.Password
	}
	db.connection.Save(&user)
	return user
}

func (db userConnection) VerifyCredential(email string, password string) interface{} {
	var user entity.User
	res := db.connection.Where("email = ?", email).Take(&user)
	if res.Error == nil {
		return user
	}

	return nil
}

func (db userConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user entity.User
	return db.connection.Where("email = ?", email).Take(&user)
}

func (db userConnection) FindUserByEmail(email string) entity.User {
	var user entity.User
	db.connection.Where("email = ?", email).Take(&user)
	return user
}

func (db userConnection) IsUserExists(userid string) (tx *gorm.DB) {
	var user entity.User
	return db.connection.Where("id = ?", userid).Take(&user)
}


func hashAndSalt(password []byte) string {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash password " + err.Error())
	}

	return string(hash)
}
