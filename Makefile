

fire: ## Build binary
	go run ./deploy/deploy.go

generate:
	solc --optimize  --abi ./contracts/Box.sol -o build
	solc --optimize --bin ./contracts/Box.sol -o build
	abigen --abi=./build/box.abi --bin=./build/box.bin --pkg=api --out=./api/box.go
