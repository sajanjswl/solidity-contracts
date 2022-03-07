# abigen --sol ./contracts/Box.sol --pkg api --out ./box/box.go 

# #  abigen --abi ./build/box.abi  --bin ./build/box.bin


# #  abigen --abi ./build/box.abi --pkg box --out ./box/box.go --bin ./build/box.bin

# solc ./contracts/Box.sol --bin --abi --optimize -o ./build/

# solc --optimize  --abi ./contracts/Box.sol -o build

# solc --optimize --bin ./contracts/Box.sol -o build

abigen --abi=./build/box.abi --bin=./build/box.bin --pkg=api --out=./api/box.go