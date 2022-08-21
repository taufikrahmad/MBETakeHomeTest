package controller

import (
	"MBETakeHomeTest/dto"
	"MBETakeHomeTest/helper"
	"MBETakeHomeTest/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt"
	"log"
	"net/http"
	"time"
)

type UserController interface {
	Update(ctx *gin.Context)
	Profile(ctx *gin.Context)
	Balance(ctx *gin.Context)
	Logout(ctx *gin.Context)
}

type userController struct {
	userService services.UserService
	userBalanceService services.UserBalanceService
	jwtService  services.JWTService
}

func (c userController) Update(ctx *gin.Context) {
	var userUpdateDTO dto.UserUpdateDTO

	errDTO := ctx.ShouldBind(&userUpdateDTO)
	log.Println(userUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}

	claims := token.Claims.(jwt.MapClaims)
	id, _ := claims["user_id"].(string)
	userUpdateDTO.ID,_ = uuid.FromString(id)
	userUpdated := c.userService.Update(userUpdateDTO)
	res := helper.BuildResponse(true,"Ok",userUpdated)
	ctx.JSON(http.StatusOK,res)
}

func (c userController) Profile(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	log.Println("header " +authHeader)
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id, _ := claims["user_id"].(string)
	userProfile := c.userService.Profile(id)
	res := helper.BuildResponse(true,"Ok !",userProfile)
	ctx.JSON(http.StatusOK,res)
}

func (c userController) Balance(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	log.Println("header " +authHeader)
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id, _ := claims["user_id"].(string)
	log.Println(id)
	userProfile := c.userBalanceService.UserCurrentBalance(id)
	res := helper.BuildResponse(true,"Ok !",userProfile)
	ctx.JSON(http.StatusOK,res)
}

func (c userController) Logout(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	log.Println("header " +authHeader)
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now()
	claims["exp"] = time.Now()
	id, _ := claims["user_id"].(string)
	userProfile := c.userService.Profile(id)
	res := helper.BuildResponse(true,fmt.Sprintf("Goodbye, %s!",userProfile.Name),nil)
	ctx.JSON(http.StatusOK,res)
}

func NewUserController(userService services.UserService,userBalanceService services.UserBalanceService, jwtService services.JWTService) UserController {
	return &userController{
		userService: userService,
		userBalanceService: userBalanceService,
		jwtService:  jwtService,
	}

}
