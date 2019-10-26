package main

import (
	"time"
)

type Truck struct {
	TruckId            string ``
	OwnershipType      string ``
	TotalNormalWeight  uint64 ``
	TotalFragileWeight uint64 ``
	ContainersAlloted  uint64
	ContainersLoaded   []Container ``
	Shedule            time.Time   ``
	Route              Route
}
