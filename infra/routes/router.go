package routes

import (
	"github.com/TulioGuaraldoB/q2-payer-challenge/infra/db"
	"github.com/TulioGuaraldoB/q2-payer-challenge/util/health"
	"github.com/TulioGuaraldoB/q2-payer-challenge/web/business"
	"github.com/TulioGuaraldoB/q2-payer-challenge/web/controller"
	"github.com/TulioGuaraldoB/q2-payer-challenge/web/repository"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes() *fiber.App {
	// Services-Settings
	db := db.OpenConnection()

	// Repositories
	userRepository := repository.NewUserRepository(db)
	walletRepository := repository.NewWalletRepository(db)

	// Businesses
	userBusiness := business.NewUserBusiness(userRepository)
	walletBusiness := business.NewWalletBusiness(walletRepository, userRepository)

	// Controllers
	userController := controller.NewUserController(userBusiness)
	walletController := controller.NewWalletController(walletBusiness)

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
	wallet.Post("", walletController.CreateWallet)
	wallet.Delete(":id", walletController.DeleteWallet)

	return app
}
