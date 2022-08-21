package controller

import (
	"MBETakeHomeTest/dto"
	"MBETakeHomeTest/entity"
	"MBETakeHomeTest/helper"
	"MBETakeHomeTest/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt"
	"log"
	"net/http"
)

type TransactionController interface {
	TopUp(ctx *gin.Context)
	WithDraw(ctx *gin.Context)
	Transfer(ctx *gin.Context)
	History(ctx *gin.Context)
}

type transactionController struct {
	authService        services.AuthService
	jwtService         services.JWTService
	transactionService services.TransactionService
	userBalanceService services.UserBalanceService
}

func (c transactionController) TopUp(ctx *gin.Context) {
	var transactionDTO dto.TransactionDTO

	errDTO := ctx.ShouldBind(&transactionDTO)
	log.Println(transactionDTO)
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
	transactionDTO.UserIDFrom, _ = uuid.FromString(id)
	userUpdated := c.transactionService.CreateTransaction(transactionDTO)

	userBalance := c.userBalanceService.UserCurrentBalance(id)
	userBalance.Balance += transactionDTO.Balance

	c.userBalanceService.UpdateUserBalance(dto.UserBalanceUpdateDTO{
		ID:      userBalance.ID,
		UserID:  transactionDTO.UserIDFrom,
		Balance: userBalance.Balance,
	})
	res := helper.BuildResponse(true, fmt.Sprintf("Topup success! Your balance is %f",userBalance.Balance), userUpdated)
	ctx.JSON(http.StatusOK, res)
}

func (c transactionController) WithDraw(ctx *gin.Context) {
	var transactionDTO dto.TransactionDTO

	errDTO := ctx.ShouldBind(&transactionDTO)
	log.Println(transactionDTO)
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


	if c.transactionService.IsEnoughBalance(id,transactionDTO.Balance) {
		response := helper.BuildErrorResponse("Failed to process", "Not Enough Balance !", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	transactionDTO.UserIDFrom, _ = uuid.FromString(id)
	transactionCreated := c.transactionService.CreateTransaction(transactionDTO)
	userBalance := c.userBalanceService.UserCurrentBalance(id)
	userBalance.Balance -= transactionDTO.Balance
	c.userBalanceService.UpdateUserBalance(dto.UserBalanceUpdateDTO{
		ID:      userBalance.ID,
		UserID:  transactionDTO.UserIDFrom,
		Balance: userBalance.Balance,
	})
	res := helper.BuildResponse(true, fmt.Sprintf("Withdraw success! Your balance is %f",userBalance.Balance), transactionCreated)
	ctx.JSON(http.StatusOK, res)
}

func (c transactionController) Transfer(ctx *gin.Context) {
	var transactionDTO dto.TransactionDTO

	errDTO := ctx.ShouldBind(&transactionDTO)
	log.Println(transactionDTO)
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


	if c.transactionService.IsEnoughBalance(id,transactionDTO.Balance) {
		response := helper.BuildErrorResponse("Failed to process", "Not Enough Balance !", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if c.transactionService.IsDestinationExist((transactionDTO.UserIDTo).String()) {
		response := helper.BuildErrorResponse("Failed to process", "Destination Not Found !", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	transactionDTO.UserIDFrom, _ = uuid.FromString(id)
	transactionCreated := c.transactionService.CreateTransaction(transactionDTO)
	userBalance := c.userBalanceService.UserCurrentBalance(id)
	userBalance.Balance -= transactionDTO.Balance
	c.userBalanceService.UpdateUserBalance(dto.UserBalanceUpdateDTO{
		ID:      userBalance.ID,
		UserID:  transactionDTO.UserIDFrom,
		Balance: userBalance.Balance,
	})

	userDestination := c.userBalanceService.UserCurrentBalance((transactionDTO.UserIDTo).String())
	log.Println("destination balance : ")
	log.Println(userDestination)
	userDestination.Balance += transactionDTO.Balance
	c.userBalanceService.UpdateUserBalance(dto.UserBalanceUpdateDTO{
		ID:      userDestination.ID,
		UserID:  transactionDTO.UserIDTo,
		Balance: userDestination.Balance,
	})
	userBalance = c.userBalanceService.UserCurrentBalance(id)
	res := helper.BuildResponse(true, fmt.Sprintf("Withdraw success! Your balance is %f",userBalance.Balance), transactionCreated)
	ctx.JSON(http.StatusOK, res)
}

func (c transactionController) History(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}

	claims := token.Claims.(jwt.MapClaims)
	id, _ := claims["user_id"].(string)
	var transactions []entity.TransactionHistory = c.transactionService.UserTransactions(id)
	res := helper.BuildResponse(true, "OK", transactions)
	ctx.JSON(http.StatusOK, res)
}

func NewTransactionController(authService services.AuthService, jwtService services.JWTService, transactionService services.TransactionService, userBalanceService services.UserBalanceService) TransactionController {
	return &transactionController{
		authService:        authService,
		jwtService:         jwtService,
		transactionService: transactionService,
		userBalanceService: userBalanceService,
	}
}
