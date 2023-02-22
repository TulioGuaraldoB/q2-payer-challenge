mysql:
	@docker-compose up -d

mysql-down:
	@docker-compose down

run:
	@echo Running...
	@go run main.go

install:
	@echo Downloading dependencies...
	@go get
	@echo Validating dependencies...
	@go mod tidy

vendor:
	@echo Generating vendor from dependencies...
	@go mod vendor

mock:
	@echo Generating mocks...
	@echo Mocking repositories...
	@mockgen -source=web/repository/userRepository.go -destination=web/repository/mock/mockUserRepository.go -package=mock
	@mockgen -source=web/repository/walletRepository.go -destination=web/repository/mock/mockWalletRepository.go -package=mock
	@mockgen -source=web/repository/transactionRepository.go -destination=web/repository/mock/mockTransactionRepository.go -package=mock
	@echo Mocking businesses...
	@mockgen -source=web/business/userBusiness.go -destination=web/business/mock/mockUserBusiness.go -package=mock
	@mockgen -source=web/business/walletBusiness.go -destination=web/business/mock/mockWalletBusiness.go -package=mock
	@mockgen -source=web/business/transactionBusiness.go -destination=web/business/mock/mockTransactionBusiness.go -package=mock
	@echo Mocking services...
	@mockgen -source=web/service/authorizationService.go -destination=web/service/mock/mockAuthorizationService.go -package=mock

test:
	@echo Running tests...
	@go test -v ./...
	@echo Tests passed!

cover:
	@echo Running test coverage...
	@go test -v ./... -coverprofile=coverage/cover.out
	@go tool cover -html=coverage/cover.out -o coverage/cover.html
	@echo Test coverage successfully completed!

swagger:
	@echo Generating documentation...
	@swag init
	@echo Swagger documentation generated successfully!