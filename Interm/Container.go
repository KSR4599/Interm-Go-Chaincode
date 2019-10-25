package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type Container struct {
	containerId   string ``
	normalWeight  uint64 `50`
	fragileWeight uint64 `0`
	allShipments  []Shipment
	route         Route  ``
	truck         Truck  ``
	readyToLoad   bool   `true`
	status        string ``
}

type Shipment struct {
	weight       uint64 `0`
	ShipmentType string ``
}

type Route struct {
	origin      string ``
	destination string ``
	DateTime    string `time.Now()`
}

func (IntermChaincode *IntermChaincode) createContainer(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	fmt.Println("The args received in the createContainer function :-", args)
	var route Route

	route.origin = string(args[1])
	route.destination = string(args[2])
	currentTime := time.Now()
	route.DateTime = currentTime.String()

	var container Container
	container.containerId = string(args[0]) + strconv.Itoa(r.Intn(999999))
	container.route = route
	container.allShipments = []Shipment{}
	container.readyToLoad = false
	container.normalWeight = 0
	container.fragileWeight = 0
	container.status = "Intransit"

	fmt.Println("Create the Container:-")
	fmt.Printf("%+v", container)
	jsonBlob, _ := json.Marshal(container)
	fmt.Println("The marhsal format of Container :-", jsonBlob)

	var cont Container
	err := json.Unmarshal(jsonBlob, &cont)

	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("The unmarshall of Container :-")
	fmt.Printf("%+v", cont)

	stub.PutState(container.containerId, jsonBlob)

	return shim.Success([]byte("successful"))
}

func (IntermChaincode *IntermChaincode) getContainer(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	fmt.Println("The arg we recieved is :-", args[0])
	bytes, _ := stub.GetState(args[0])
	if bytes == nil {
		return shim.Error("Not Container is Found")
	}

	var container Container
	_ = json.Unmarshal(bytes, &container)

	fmt.Println("The container we got", container)
	return shim.Success([]byte("success"))
}
