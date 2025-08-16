package main

type ParkingSpot struct {
	SpotId        int
	VehicleType   VehicleType
	parkedVehicle *Vehicle
}

func NewParkingSpot(id int, vehicleType VehicleType) *ParkingSpot {
	return &ParkingSpot{SpotId: id, VehicleType: vehicleType}
}

func (s *ParkingSpot) IsAvailable() bool {
	return s.parkedVehicle == nil
}

func (s *ParkingSpot) CanParkVehicle(vehicleType VehicleType) bool {
	return s.IsAvailable() && vehicleType == s.VehicleType
}

func (s *ParkingSpot) ParkVehicle(vehicle *Vehicle) bool {
	if s.IsAvailable() && vehicle.VehicleType == s.VehicleType {
		s.parkedVehicle = vehicle
		return true
	}
	return false
}

func (s *ParkingSpot) UnParkVehicle() *Vehicle {
	v := s.parkedVehicle
	s.parkedVehicle = nil
	return v
}
