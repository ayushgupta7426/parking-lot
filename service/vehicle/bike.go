package vehicle

import (
	"time"
)

type Bike struct{
	Vehicle
}

func NewBike(vehicleNumber string) IVehicle {
	return &Bike{
		Vehicle: Vehicle{
		Id: vehicleNumber,
		VehicleType: BIKE,
		EntryTime: time.Now(),
		},
	}
}

func (b *Bike)GetVehicle() *Vehicle {
	return &b.Vehicle
}
