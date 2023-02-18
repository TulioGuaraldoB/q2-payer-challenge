package routes

import (
	"net/http"

	"github.com/TulioGuaraldoB/q2-payer-challenge/infra/db"
	"github.com/TulioGuaraldoB/q2-payer-challenge/util/health"
	"github.com/TulioGuaraldoB/q2-payer-challenge/web/business"
	"github.com/TulioGuaraldoB/q2-payer-challenge/web/controller"
	"github.com/TulioGuaraldoB/q2-payer-challenge/web/repository"
	"github.com/TulioGuaraldoB/q2-payer-challenge/web/service"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes() *fiber.App {
	// Services-Settings
	db := db.OpenConnection()
	httpClient := &http.Client{}

	// Services
	authorizationService := service.NewAuthorizationService(httpClient)

	// Repositories
	userRepository := repository.NewUserRepository(db)
	walletRepository := repository.NewWalletRepository(db)
	transactionRepository := repository.NewTransactionRepository(db)

	// Businesses
	userBusiness := business.NewUserBusiness(userRepository)
	walletBusiness := business.NewWalletBusiness(walletRepository, userRepository, authorizationService)
	transactionBusiness := business.NewTransactionBusiness(transactionRepository)

	// Controllers
	userController := controller.NewUserController(userBusiness)
	walletController := controller.NewWalletController(walletBusiness)
	transactionController := controller.NewTransactionController(transactionBusiness)

	app := fiber.New()
	app.Get("health", health.HealthCheck)

	api := app.Group("api")
	v1 := api.Group("v1")

	user := v1.Group("user")
	user.Get("", userController.GetAllUsers)
	user.Get(":id", userController.GetUserById)
	user.Get("email/:email", userController.GetUserByEmail)
	user.Get("cpf/:cpf", userController.GetUserByCpf)
	user.Post("login", userController.GetUserByCredentials)
	user.Post("", userController.CreateUser)
	user.Delete("", userController.DeleteUser)

	wallet := v1.Group("wallet")
	wallet.Post("user", walletController.GetWalletByUserCredentials)
	wallet.Post("deposit", walletController.DepositToWalletBalance)
	wallet.Post("", walletController.CreateWallet)
	wallet.Delete(":id", walletController.DeleteWallet)

	transaction := v1.Group("transaction")
	transaction.Get(":id", transactionController.GetTransactionById)

	return app
}
