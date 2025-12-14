package vehicle

import (
	"time"
)

type Car struct {
	Vehicle
}

func NewCar(vehicleNumber string) IVehicle {
	return &Bike{
		Vehicle: Vehicle{
			Id:            vehicleNumber,
			VehicleType:      CAR,
			EntryTime:     time.Now(),
		},
	}
}

func (b *Car) GetVehicle() *Vehicle {
	return &b.Vehicle
}
