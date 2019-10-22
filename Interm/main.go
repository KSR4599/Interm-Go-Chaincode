package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type IntermChaincode struct {
}

func (token *ERC20TokenChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {

	fmt.Println("Init executed")
	_, args := stub.GetFunctionAndParameters()

	return shim.Success([]byte(jsonERC20))
}

// Invoke method
func (token *ERC20TokenChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {

	// Get the function name and parameters
	function, args := stub.GetFunctionAndParameters()

	fmt.Println("Invoke executed : ", function, " args=", args)

	switch {

	// Query function
	case function == "createContainer":
		return *ERC20TokenChaincode.createContainer(stub, args)
	case function == "clearContainer":
		return *ERC20TokenChaincode.clearContainer(stub, args)
	case function == "loadContainer":
		return *ERC20TokenChaincode.loadContainer(stub, args)
	case function == "getContainer":
		return *ERC20TokenChaincode.getContainer(stub, args)
	case function == "readyContainer":
		return *ERC20TokenChaincode.readyContainer(stub, args)
	case function == "assignTruck":
		return *ERC20TokenChaincode.assignTruck(stub, args)
	case function == "addTruck":
		return *ERC20TokenChaincode.addTruck(stub, args)

	}

	return errorResponse("Invalid function", 1)
}

func main() {
	fmt.Println("Started the Interm Chaincode")
	err := shim.Start(new(IntermChaincode))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}
