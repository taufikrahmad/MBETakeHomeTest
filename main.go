package main

import (
	"MBETakeHomeTest/config"
	"MBETakeHomeTest/controller"
	"MBETakeHomeTest/middleware"
	"MBETakeHomeTest/repository"
	"MBETakeHomeTest/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db                        *gorm.DB                         = config.SetupDatabaseConnection()
	userRepository            repository.UserRepository        = repository.NewUserRepository(db)
	userBalanceRepository     repository.UserBalanceRepository = repository.NewUserBalanceRepository(db)
	transactionRepository     repository.TransactionRepository = repository.NewUserTransactionRepository(db)
	jwtService                services.JWTService              = services.NewJWTService()
	authService               services.AuthService             = services.NewAuthService(userRepository)
	userService               services.UserService             = services.NewUserService(userRepository)
	userBalanceService        services.UserBalanceService      = services.NewUserBalanceService(userBalanceRepository)
	transactionBalanceService services.TransactionService      = services.NewTransactionService(transactionRepository, userBalanceRepository, userRepository)
	authController            controller.AuthController        = controller.NewAuthController(authService, jwtService, userBalanceService)
	userController            controller.UserController        = controller.NewUserController(userService, userBalanceService, jwtService)
	transactionController     controller.TransactionController = controller.NewTransactionController(authService, jwtService, transactionBalanceService,userBalanceService)
)

func main() {
	r := gin.Default()
	//,middleware.AuthorizeJWT(jwtService)
	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.GET("/balance", userController.Balance)
		userRoutes.GET("/logout", userController.Logout)
		userRoutes.PUT("/", userController.Update)
	}

	transactionRoutes := r.Group("api/transaction", middleware.AuthorizeJWT(jwtService))
	{
		transactionRoutes.POST("/topup", transactionController.TopUp)
		transactionRoutes.POST("/withdraw", transactionController.WithDraw)
		transactionRoutes.POST("/transfer", transactionController.Transfer)
		transactionRoutes.GET("/", transactionController.History)
	}

	r.Run(":5000")
}
