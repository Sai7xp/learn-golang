package main

type ParkingFloor struct {
	// each level has list of spots
	floor            int
	ParkingSpotsList []*ParkingSpot
}

func NewParkingFloor(capacity, floor int) *ParkingFloor {
	newLevel := &ParkingFloor{floor: floor}
	spots := make([]*ParkingSpot, capacity)
	for i := range capacity {
		if i&1 == 1 {
			spots[i] = NewParkingSpot(i, Car)
		} else {
			spots[i] = NewParkingSpot(i, Bike)
		}
	}
	newLevel.ParkingSpotsList = spots
	return newLevel
}

func (l *ParkingFloor) FindAvailableSpot(vehicleType VehicleType) *ParkingSpot {
	for _, eachSpot := range l.ParkingSpotsList {
		if eachSpot.CanParkVehicle(vehicleType) {
			return eachSpot
		}
	}
	return nil
}
