package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type Truck struct {
	TruckId            string ``
	OwnershipType      string ``
	TotalNormalWeight  uint64 ``
	TotalFragileWeight uint64 ``
	ContainersAlloted  uint64
	ContainersLoaded   []Container ``
	Schedule           string      ``
	Route              Route
}

func (IntermChaincode *IntermChaincode) createTruck(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	var truck Truck
	truck.TruckId = string(args[0])
	truck.OwnershipType = string(args[1])
	truck.TotalFragileWeight, _ = strconv.ParseUint(args[2], 10, 64)
	truck.TotalNormalWeight, _ = strconv.ParseUint(args[3], 10, 64)
	truck.ContainersAlloted = 0
	truck.ContainersLoaded = []Container{}
	currentTime := time.Now()
	truck.Schedule = currentTime.String()

	var route Route

	route.Origin = string(args[4])
	route.Destination = string(args[5])
	currentTime = time.Now()
	route.DateTime = currentTime.String()
	truck.Route = route

	fmt.Println("The following truck got created :-", truck)

	jsonBlob, _ := json.Marshal(truck)

	stub.PutState(truck.TruckId, jsonBlob)

	return shim.Success([]byte(" Truck Creation successful"))
}

func (IntermChaincode *IntermChaincode) getTruck(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	fmt.Println("The arg we recieved in getTruck is :-", args[0])
	jsonBlob, _ := stub.GetState(args[0])
	if jsonBlob == nil {
		return shim.Error("No Truck is Found")
	}

	var tru Truck

	err := json.Unmarshal(jsonBlob, &tru)

	if err != nil {
		fmt.Println("error:", err)
	}

	var truck1 Truck
	_ = json.Unmarshal(jsonBlob, &truck1)

	fmt.Println("The truck we got :-")
	fmt.Printf("%+v", truck1)
	return shim.Success([]byte("successfully Got the Truck"))
}

func (IntermChaincode *IntermChaincode) assignTruck(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	jsonBlob, _ := stub.GetState(args[0])
	if jsonBlob == nil {
		return shim.Error("No Container is Found")
	}

	var cont Container

	err := json.Unmarshal(jsonBlob, &cont)

	if err != nil {
		fmt.Println("error:", err)
	}

	var container Container
	_ = json.Unmarshal(jsonBlob, &container)

	fmt.Println("The container we got :-")
	fmt.Printf("%+v", container)

	if container.ReadyToLoad == false {
		shim.Error("The container is not yet ready to load")
	} else {

		jsonBlob1, _ := stub.GetState(args[1])
		var tru Truck

		_ = json.Unmarshal(jsonBlob1, &tru)

		tru.TotalFragileWeight = tru.TotalFragileWeight + container.FragileWeight
		tru.TotalNormalWeight = tru.TotalNormalWeight + container.NormalWeight
		tru.ContainersAlloted = tru.ContainersAlloted + 1

		tru.ContainersLoaded = append(tru.ContainersLoaded, container)
		fmt.Println("The updated truck :-")
		fmt.Printf("%+v", tru)

		jsonBlob, _ = json.Marshal(tru)

		stub.PutState(tru.TruckId, jsonBlob)

		return shim.Success([]byte(" Truck Assignment successful"))
	}
}
