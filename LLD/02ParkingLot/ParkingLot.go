package main

import "sync"

// main parking manager, all interactions will be handled by this guy
type ParkingLot struct {
	// multiple levels
	levels            []*ParkingFloor
	activeSpots       map[string]*ParkingSpot
	totalCollectedFee float32
}

var instance *ParkingLot
var once sync.Once

func GetParkingLotInstance() *ParkingLot {
	once.Do(func() {
		instance = &ParkingLot{
			activeSpots: make(map[string]*ParkingSpot),
		}
	})
	return instance
}

func (p *ParkingLot) AddLevel(spotsCapacity int) {
	newLevel := NewParkingFloor(spotsCapacity, len(p.levels)+1)
	p.levels = append(p.levels, newLevel)
}

func (p *ParkingLot) ParkVehicle(vehicle *Vehicle) bool {
	// first get the availability spot and then park the vehicle in that spot
	var parkingSpot *ParkingSpot
	for _, eachLevel := range p.levels {
		if s := eachLevel.FindAvailableSpot(vehicle.VehicleType); s != nil {
			parkingSpot = s
			break
		}
	}
	if parkingSpot == nil {
		return false
	}
	parkingSpot.ParkVehicle(vehicle)
	p.activeSpots[vehicle.LicenseNum] = parkingSpot
	return true
}

func (p *ParkingLot) UnParkVehicle(licenseNum string) *Vehicle {
	parkedSpot := p.activeSpots[licenseNum]
	if parkedSpot == nil {
		return nil
	}
	parkedVehicle := parkedSpot.UnParkVehicle()
	p.totalCollectedFee += 10
	delete(p.activeSpots, licenseNum)
	return parkedVehicle
}

func (p *ParkingLot) GetTotalCollectedFee() float32 {
	return p.totalCollectedFee
}
