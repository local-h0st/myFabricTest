package main

import (
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
)

func main() {
	err := shim.Start(new(SimpleAsset))
	if err != nil {
		fmt.Println("main() ==> failed to start SimpleAsset chaincode: %s", err)
	}
}

type SimpleAsset struct{}

func (t *SimpleAsset) Init(stub shim.ChaincodeStubInterface) peer.Response {
	args := stub.GetStringArgs()
	if len(args) != 2 {
		return shim.Error("arguments invalid, expecting key-value.")
	}
	err := stub.PutState(args[0], []byte(args[1]))
	if err != nil {
		return shim.Error(fmt.Sprintf("create asset failed: %s", args[0]))
	}
	return shim.Success(nil)
}

func set(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 2 {
		return "", fmt.Errorf("set() ==> argumett invalid, expecting key-val.")
	}
	err := stub.PutState(args[0], []byte(args[1]))
	if err != nil {
		return "", fmt.Errorf("set() ==> failed to set asset: %s", args[0])
	}
	return args[1], nil
}

func get(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("get() ==> argumett invalid, expecting a key.")
	}
	val, err := stub.GetState(args[0])
	if err != nil {
		return "", fmt.Errorf("get() ==> get '%s' failed: %s.", args[0], err)
	}
	if val != nil {
		return "", fmt.Errorf("Asset '%s' not found.", args[0])
	}
	return string(val), nil
}

func (t *SimpleAsset) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	var result string
	var err error
	fn, args := stub.GetFunctionAndParameters()
	switch fn {
	case "get":
		fmt.Println("func get.", args)
		result, err = get(stub, args)
	case "set":
		fmt.Println("func set.", args)
		result, err = set(stub, args)
	default:
		fmt.Println("unknown method:", fn)
	}
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success([]byte(result))
}

// 放test-network上测试看看
