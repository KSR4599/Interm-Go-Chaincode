package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type IntermChaincode struct {
}

func (token *IntermChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {

	fmt.Println("Init executed")
	_, args := stub.GetFunctionAndParameters()

	return shim.Success([]byte(jsonERC20))
}

// Invoke method
func (token *IntermChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {

	// Get the function name and parameters
	function, args := stub.GetFunctionAndParameters()

	fmt.Println("Invoke executed : ", function, " args=", args)

	switch {

	// Query function
	case function == "createContainer":
		return *IntermChaincode.createContainer(stub, args)
	case function == "clearContainer":
		return *IntermChaincode.clearContainer(stub, args)
	case function == "loadContainer":
		return *IntermChaincode.loadContainer(stub, args)
	case function == "getContainer":
		return *IntermChaincode.getContainer(stub, args)
	case function == "readyContainer":
		return *IntermChaincode.readyContainer(stub, args)
	case function == "assignTruck":
		return *IntermChaincode.assignTruck(stub, args)
	case function == "addTruck":
		return *IntermChaincode.addTruck(stub, args)

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
