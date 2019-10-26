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
	ContainerId   string     `json:"containerId"`
	NormalWeight  uint64     `json:"normalWeight"`
	FragileWeight uint64     `json:"fragileWeight"`
	AllShipments  []Shipment `json:"allShipments"`
	Route         Route      `json:"route"`
	Truck         Truck      `json:"truck"`
	ReadyToLoad   bool       `json:"readyToLoad"`
	Status        string     `json:"status"`
}

type Shipment struct {
	Weight       uint64 `0`
	ShipmentType string ``
}

type Route struct {
	Origin      string ``
	Destination string ``
	DateTime    string `time.Now()`
}

func (IntermChaincode *IntermChaincode) createContainer(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	var route Route

	route.Origin = string(args[1])
	route.Destination = string(args[2])
	currentTime := time.Now()
	route.DateTime = currentTime.String()

	var container Container
	container.ContainerId = string(args[0]) + strconv.Itoa(r.Intn(999999))
	container.Route = route
	container.AllShipments = []Shipment{}
	container.ReadyToLoad = false
	container.NormalWeight = 0
	container.FragileWeight = 0
	container.Status = "Intransit"

	fmt.Println("The following container got created :-", container)
	jsonBlob, _ := json.Marshal(container)

	stub.PutState(container.ContainerId, jsonBlob)

	return shim.Success([]byte("Container Creation successful"))
}

func (IntermChaincode *IntermChaincode) getContainer(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	fmt.Println("The arg we recieved in getContainer is :-", args[0])
	jsonBlob, _ := stub.GetState(args[0])
	if jsonBlob == nil {
		return shim.Error("No Container is Found")
	}

	var cont Container

	err := json.Unmarshal(jsonBlob, &cont)

	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("The container we got :-")
	fmt.Printf("%+v", cont)
	return shim.Success([]byte("successfully got the container"))
}

func (IntermChaincode *IntermChaincode) loadContainer(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	jsonCont, _ := stub.GetState(args[0])
	if jsonCont == nil {
		return shim.Error("No Container is Found")
	}

	var contt Container

	err := json.Unmarshal(jsonCont, &contt)

	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("The container we got :-")
	fmt.Printf("%+v", contt)

	if contt.ReadyToLoad == true {
		shim.Error("The container is already ready to load into Truck. Can't add load anymore.")
	} else {
		fmt.Println("The container status is not ready to load. So we are good :-)")
	}

	var weight, _ = strconv.ParseUint(args[2], 10, 64)

	if args[1] == "Fragile" {

		if contt.FragileWeight+weight > 400 {
			shim.Error("Container will be overloaded!")
		} else {
			contt.FragileWeight = contt.FragileWeight + weight
		}
	} else {
		if contt.NormalWeight+weight > 600 {
			shim.Error("Container will be overloaded!")
		} else {
			contt.NormalWeight = contt.NormalWeight + weight
		}
	}

	var shipment Shipment
	shipment.Weight, _ = strconv.ParseUint(args[2], 10, 64)
	shipment.ShipmentType = args[1]

	contt.AllShipments = append(contt.AllShipments, shipment)

	fmt.Println("The Updated Container :-", contt)

	jsonBlob, _ := json.Marshal(contt)

	stub.PutState(contt.ContainerId, jsonBlob)

	return shim.Success([]byte("Container Loading successful"))
}
