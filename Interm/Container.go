package main

import (
	"math/rand"
	"strconv"
	"time"
	"github.com/hyperledger/fabric/core/chaincode/lib/cid"
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
	origin      string ``
	destination string ``
	DateTime    string `time.Now()`
}

func createContainer(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	container = Container{}
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	contid, _, _ := cid.GetAttributeValue(stub, "containerId")
	origin, _, _ := cid.GetAttributeValue(stub, "origin")
	destination, _, _ := cid.GetAttributeValue(stub, "destination")
	conatiner.containerId := contid + strconv.Itoa(r.Intn(999999))
	route := Route{}
	route.origin := origin
	route.destination := destination
	route.DateTime := time.Now()
	container.route := route
	container.allShipments = [];
	container.readyToLoad = false;
	container.normalWeight = 0;
    container.fragileWeight = 0;
	container.status = "Intransit";
	jsonContainer, _ := json.Marshal(container)
	stub.PutState(containter.containerId, jsonContainer)
	fmt.Println("Create the Container:-", container)
	
	return shim.Success([]byte(container))
}

func getContainer(stub shim.ChaincodeStubInterface, args []string) peer.Response { 

	qry := `{
		"selector": {
		   "containerId": {
			  `args[0]`
		   }
		}
	 }`

	 QryContainer, err := stub.GetQueryResult(qry)

	 if err != nil {
		return shim.Error("Error in fetching the desired container !!!! "+err.Error())
	}
	return shim.Success([]byte(QryContainer))
}
