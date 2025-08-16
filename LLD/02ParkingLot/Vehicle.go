package main

type VehicleType int

const (
	Truck VehicleType = iota + 1
	Bike
	Car
)

type Vehicle struct {
	LicenseNum  string // license num
	VehicleType VehicleType
}

func NewVehicle(licenseNum string, vType VehicleType) *Vehicle {
	return &Vehicle{
		LicenseNum:  licenseNum,
		VehicleType: vType,
	}
}
