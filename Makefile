dependency:
	@echo ">> Downloading Dependencies"
	@go mod download

swag-init:
	@echo ">> Running swagger init"
	@swag init

run-api: dependency swag-init
	@echo ">> Running API Server"
	@go run main.go server-http

run-grpc: dependency
	@echo ">> Running gRPC Server"
	@go run main.go serve-grpc

generate-abi:
	@echo  ">> Generating ABI"
	@solcjs --abi ./pkg/contracts/Votechain.sol -o ./pkg/binding

generate-gobind:
	@echo ">> Generating Go Bindings"
	@abigen --abi ./pkg/binding/Votechain.abi --pkg binding --type Votechain --out ./pkg/binding/Votechain.go

change-abi:
	@mv ./pkg/binding/pkg_contracts_Votechain_sol_Votechain.abi ./pkg/binding/Votechain.abi

migrate-up:
	@echo ">> Running Migrate Up"
	@migrate -path db/migrations -database "postgres://postgres:1235813@localhost:5433/ums?sslmode=disable" up

migrate-down:
	@echo ">> Running Migrate down"
	@migrate -path db/migrations -database "postgres://postgres:1235813@localhost:5433/ums?sslmode=disable" down

remock:
	#https://github.com/vektra/mockery
	@echo ">> Mock Repositories"
	@mockery --all --recursive --dir ./internal/domain/repository --output ./internal/domain/repository/mocks_repository --outpkg mocks_repository

	@echo ">> Mock Usecases"
	@mockery --all --dir ./internal/usecases --output ./internal/usecases/mocks_usecases --outpkg mocks_usecases

	@echo ">> Mock Interfaces"
	@mockery --all --recursive --dir ./internal/interfaces --output ./internal/interfaces/mocks_interfaces --outpkg mocks_interfaces

	@echo ">> Mock Infra"
	@mockery --all --recursive --dir ./internal/infrastructures --output ./internal/infrastructures/mocks_infrastructures --outpkg mocks_infrastructures