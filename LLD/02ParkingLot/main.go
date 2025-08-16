package main

import "fmt"

/*
Functional Requirements
- Multi Level Parking Lot
- What type of vehicles
- Park Vehicle & Unpark Vehicle (occupy & free up the space functionality)
- Check for parking space in all Levels, consider vehicle type as well
- Payment ?

Entities
- ParkingFloor : floorNumber,
- Parking Spot : spotId, vehicleType, *Vehicle
- enum VehicleType : Car, Bike, Truck
- Vehicle - VehicleType, licenseNumber
*/

func main() {
	fmt.Println("Parking Lot - LLD")
	newparkingLot := GetParkingLotInstance()
	newparkingLot.AddLevel(0)

	myCar := NewVehicle("76", Car)
	if ok := newparkingLot.ParkVehicle(myCar); ok {
		fmt.Printf("%s vehicle parked\n", myCar.LicenseNum)
	} else {
		fmt.Println("No parking space available. Sorry")
		return
	}
	if v := newparkingLot.UnParkVehicle(myCar.LicenseNum); v != nil {
		fmt.Printf("here is your vehicle %#v\n", v)
	} else {
		fmt.Println("Vehicle not found with the given licenseNum: ", myCar.LicenseNum)
	}
	fmt.Println("Total Fee: ", newparkingLot.GetTotalCollectedFee())
}
