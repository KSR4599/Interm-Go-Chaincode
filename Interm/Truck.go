package main

 import (

	// The shim package
	"github.com/hyperledger/fabric/core/chaincode/shim"

	// peer.Response is in the peer package
	"github.com/hyperledger/fabric/protos/peer"

	// Client Identity Library
	"github.com/hyperledger/fabric/core/chaincode/lib/cid"

	"strconv"
 )

type Truck struct {
	truckID string 'TRUCK123'
}