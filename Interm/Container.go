package main

import (
	"fmt"
)

type Container struct {
	containerId     string ``
	containerNumber string ``
	normalWeight    uint64 `50`
	fragileWeight   uint64 `0`
	allShipments    []Shipment
	route           Route  ``
	truck           string ``
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

func display() {

	container := Container{
		containerId:     "CONT123",
		containerNumber: "CON123",
		normalWeight:    10,
		fragileWeight:   20,
		readyToLoad:     false,
		status:          "Nil",
	}

	shipment := Shipment{weight: 30, ShipmentType: "Normal"}
	shipment1 := Shipment{weight: 40, ShipmentType: "Fragile"}
	container.allShipments = append(container.allShipments, shipment)
	container.allShipments = append(container.allShipments, shipment1)
	container.route.origin = "HYD"
	container.route.destination = "BAN"
	container.route.DateTime = "Now"
	fmt.Println("Container :-", container)
}

func main() {
	fmt.Println("Started....")

	display()
}
