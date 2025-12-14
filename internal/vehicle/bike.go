package vehicle

import (
)

type Bike struct{
	vehicleNumber            string
	vehicleType      VehicleType
}

func NewBike(vehicleNumber string) IVehicle {
	return &Bike{
		vehicleNumber: vehicleNumber,
		vehicleType: BikeType,
	}
}

func (c *Bike) VehicleNumber() string {
	return c.vehicleNumber
}

func (c *Bike) VehicleType() VehicleType {
	return c.vehicleType
}
