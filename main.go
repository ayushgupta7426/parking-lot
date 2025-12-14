package main

import (
	"fmt"
	parkinggate "parking/internal/entrance"
	"parking/internal/parking"
	"parking/internal/vehicle"
	"time"
)

func main() {
	parkingManager := parking.New()
	parkingManager.Add(vehicle.BikeType, 4)
	parkingManager.Add(vehicle.CarType, 4)

	parkkingGate := parkinggate.NewParkingGate(parkingManager)
	bike := vehicle.NewBike("abc")
	car := vehicle.NewCar("xyz")
	// car:=vehicle.NewCar("xyz")
	ticket, err := parkkingGate.Park(bike)
	if err != nil {
		fmt.Printf("failed to checkin, err: %v", err)
		return
	}
	fees, err := parkkingGate.Exit(ticket.TicketID, time.Now().Add(5*time.Hour))
	if err != nil {
		fmt.Printf("failed to checkout take default fees, err: %v", err)
		return
	}

	fmt.Printf("ticket: %v\n", fees)

	ticket2, err := parkkingGate.Park(car)
	if err != nil {
		fmt.Printf("failed to checkin, err: %v", err)
		return
	}
	ticket2, err = parkkingGate.Exit(ticket2.TicketID, time.Now().Add(5*time.Hour))
	if err != nil {
		fmt.Printf("failed to checkout take default fees, err: %v", err)
		return
	}
	fmt.Printf("ticket: %v\n", ticket2)
}
