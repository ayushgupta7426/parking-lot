package parking

import "parking/internal/vehicle"

type parkingSpot struct {
	id          string
	isOccupied  bool
	floor       int
	vehicleType vehicle.VehicleType
}

func (ps *parkingSpot) GetFloor() int {
	return ps.floor
}

func (ps *parkingSpot) GetVehicleType() vehicle.VehicleType {
	return ps.vehicleType
}

func (ps *parkingSpot) GetId() string {
	return ps.id
}