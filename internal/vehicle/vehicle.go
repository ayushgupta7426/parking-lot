package vehicle


type VehicleType int

const (
	CarType VehicleType = iota
	BikeType
)



type IVehicle interface {
	VehicleNumber() string
	VehicleType() VehicleType 
}
