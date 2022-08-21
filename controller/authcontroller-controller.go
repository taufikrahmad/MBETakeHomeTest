package controller

import (
	"MBETakeHomeTest/dto"
	"MBETakeHomeTest/entity"
	"MBETakeHomeTest/helper"
	"MBETakeHomeTest/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	authService services.AuthService
	jwtService  services.JWTService
	balanceService  services.UserBalanceService
}

func NewAuthController(authService services.AuthService,jwtService services.JWTService,balanceService  services.UserBalanceService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
		balanceService : balanceService,
	}
}
func (c *authController) Login(ctx *gin.Context) {
	var loginDTO dto.LoginDTO
	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	authResult := c.authService.VerifyCredential(loginDTO.Email, loginDTO.Password)
	if v, ok := authResult.(entity.User); ok {
		generateToken := c.jwtService.GenerateToken(v.ID.String())
		v.Token = generateToken
		userBalance := c.balanceService.UserCurrentBalance(v.ID.String())
		response := helper.BuildResponse(true, fmt.Sprintf("Hello, %s! Your Balance is $%f",v.Name, userBalance.Balance), v)
		ctx.JSON(http.StatusOK, response)
		return
	}

	response := helper.BuildErrorResponse("Please check your credential", "Invalid creadential", helper.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (c *authController) Register(ctx *gin.Context) {
	var registerDTO dto.RegisterDTO
	errDTO := ctx.ShouldBind(&registerDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !c.authService.IsDuplicateEmail(registerDTO.Email) {
		response := helper.BuildErrorResponse("Failed to process", "Duplicate email", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	} else {
		//create new user
		createUser := c.authService.CreateUser(registerDTO)
		//insert balance to 0
		userBalanceDTO := &dto.UserBalanceDTO{}
		userBalanceDTO.UserID = createUser.ID
		userBalanceDTO.Balance = 0
		createUserBalance := c.balanceService.InsertUserBalance(*userBalanceDTO)
		token := c.jwtService.GenerateToken(createUser.ID.String())
		createUser.Token = token
		createUser.Balance = &createUserBalance
		response := helper.BuildResponse(true, "Registration success, Hello "+createUser.Name+" !", createUser)
		ctx.JSON(http.StatusCreated, response)
		return
	}
}
