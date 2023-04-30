// 千万别部署到test_network上，会直接GG

package main

import (
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	// main() generates chaincode with contract I have written.
	fmt.Println("#Alternate World#  my chaincode 00, a test chaincode built for test-network.")
	c := new(MySimpleContarct)
	cc, err := contractapi.NewChaincode(c)
	if err != nil {
		panic(err.Error())
	}
	// start the chaincode
	err = cc.Start()
	if err != nil {
		panic(err.Error())
	}

}

type MySimpleContarct struct {
	contractapi.Contract
}

func (mySimCon *MySimpleContarct) Create(ctx contractapi.TransactionContextInterface, k string, v string) error {
	// Create() adds a new key-value to the world state.
	existing, err := ctx.GetStub().GetState(k)
	if err != nil {
		return fmt.Errorf("failed to get state of key %s.", k)
	}
	if existing != nil {
		return fmt.Errorf("failed to create, pair %s-%s already exists.", k, string(existing))
	}
	err = ctx.GetStub().PutState(k, []byte(v))
	if err != nil {
		return fmt.Errorf("failed to put state.")
	}
	return nil
}

func (mySimCon *MySimpleContarct) Update(ctx contractapi.TransactionContextInterface, k string, v string) error {
	// Update() updates the value of an existing key.
	existing, err := ctx.GetStub().GetState(k)
	if err != nil {
		return fmt.Errorf("failed to get state of key %s.", k)
	}
	if existing == nil {
		return fmt.Errorf("failed to update key %s, because it doesnt exists.", k)
	}
	err = ctx.GetStub().PutState(k, []byte(v))
	if err != nil {
		return fmt.Errorf("failed to put state.")
	}
	return nil
}

func (mySimCon *MySimpleContarct) Read(ctx contractapi.TransactionContextInterface, k string) (string, error) {
	// Read() returns the value of a key.
	existing, err := ctx.GetStub().GetState(k)
	if err != nil {
		return "", fmt.Errorf("failed to get state of key %s.", k)
	}
	if existing == nil {
		return "", fmt.Errorf("failed to read, key %s doesnt exists.", k)
	}
	return string(existing), nil
}
