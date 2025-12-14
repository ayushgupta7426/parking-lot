package vehicle

import (
)

type Car struct {
	vehicleNumber            string
	vehicleType      VehicleType
}

func NewCar(vehicleNumber string) IVehicle {
	return &Bike{
			vehicleNumber:            vehicleNumber,
			vehicleType:      CarType,
	}
}

func (c *Car) VehicleNumber() string {
	return c.vehicleNumber
}

func (c *Car) VehicleType() VehicleType {
	return c.vehicleType
}
