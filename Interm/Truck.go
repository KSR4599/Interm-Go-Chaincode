package main

import (
	"time"
)

type Truck struct {
	truckId            string ``
	ownershipType      string ``
	totalNormalWeight  uint64 ``
	totalFragileWeight uint64 ``
	containersAlloted  uint64
	containersLoaded   []Container ``
	shedule            time.Time   ``
	route              Route
}
