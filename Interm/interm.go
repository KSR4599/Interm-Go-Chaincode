package main

import (
	"fmt"

	
	"github.com/hyperledger/fabric/core/chaincode/shim"

	
	"github.com/hyperledger/fabric/protos/peer"

	
	"strconv"

	
	"encoding/json"
)

type IntermChaincode struct {
}




func main() {
	fmt.Println("Started....")
	err := shim.Start(new(IntermChaincode))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}
