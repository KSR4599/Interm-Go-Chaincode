package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/lib/cid"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type Container struct {
	containerId     string ``
	containerNumber string ``
	normalWeight    uint64 `50`
	fragileWeight   uint64 `0`
	allShipments    []Shipment
	route           Route  ``
	truck           Truck  ``
	readyToLoad     bool   `true`
	status          string ``
}

type Shipment struct {
	weight       uint64 `0`
	ShipmentType string ``
}

type Route struct {
	origin      string    ``
	destination string    ``
	DateTime    time.Time `time.Now()`
}

func (IntermChaincode *IntermChaincode) createContainer(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	contid, _, _ := cid.GetAttributeValue(stub, "containerId")
	origin1, _, _ := cid.GetAttributeValue(stub, "origin")
	destination1, _, _ := cid.GetAttributeValue(stub, "destination")
	var route Route

	route.origin = origin1
	route.destination = destination1
	route.DateTime = time.Now()

	var container Container
	container.containerId = contid + strconv.Itoa(r.Intn(999999))
	container.route = route
	container.allShipments = []Shipment{}
	container.readyToLoad = false
	container.normalWeight = 0
	container.fragileWeight = 0
	container.status = "Intransit"

	jsonContainer, _ := json.Marshal(container)
	stub.PutState(container.containerId, jsonContainer)

	fmt.Println("Create the Container:-", container)
	return shim.Success([]byte("success"))
}

func (IntermChaincode *IntermChaincode) getContainer(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	qry := `{
		"selector": {
		   "containerId": {
			  "$eq": `
	qry += args[0]
	qry += `
	  
		   }
		}
	 }`

	QryContainer, err := stub.GetQueryResult(qry)

	if err != nil {
		return shim.Error("Error in fetching the desired container !!!! " + err.Error())
	} else {
		fmt.Println("The Query Value", QryContainer)
	}
	return shim.Success([]byte("success"))
}
