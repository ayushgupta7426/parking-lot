package main

import (
	"fmt"
	parkinggate "parking/service/entrance"
	"parking/service/parking"
	"parking/service/vehicle"
	"time"
)

func main() {
	parkingManager := parking.New()
	parkingManager.Add("bike", 4)
	parkingManager.Add("car", 4)

	parkkingGate := parkinggate.NewParkingGate(parkingManager)
	bike := vehicle.NewBike("abc")
	car := vehicle.NewCar("xyz")
	// car:=vehicle.NewCar("xyz")
	err := parkkingGate.CheckIn(bike)
	if err != nil {
		fmt.Printf("failed to checkin, err: %v", err)
		return
	}
	fees, err := parkkingGate.CheckOut("abc", time.Now().Add(5*time.Hour))
	if err != nil {
		fmt.Printf("failed to checkout take default fees, err: %v", err)
		return
	}

	fmt.Printf("total fees: %v\n", fees)

	err = parkkingGate.CheckIn(car)
	if err != nil {
		fmt.Printf("failed to checkin, err: %v", err)
		return
	}
	fees1, err := parkkingGate.CheckOut("xyz", time.Now().Add(5*time.Hour))
	if err != nil {
		fmt.Printf("failed to checkout take default fees, err: %v", err)
		return
	}
	fmt.Printf("total fees: %v\n", fees1)
}
