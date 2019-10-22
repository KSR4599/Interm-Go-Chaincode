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
	fmt.Println("Create Container Called with Args of", args)
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	contid, _, _ := cid.GetAttributeValue(stub, "containerId")
	origin, _, _ := cid.GetAttributeValue(stub, "origin")
	destination, _, _ := cid.GetAttributeValue(stub, "destination")
	containerId := contid + strconv.Itoa(r.Intn(999999))
	route := Route{}
	route.origin := origin
	route.destination := destination
	route.DateTime := time.Now()
}
