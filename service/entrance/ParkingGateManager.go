package parkinggate

import (
	"fmt"
	"parking/service/parking"
	"parking/service/strategy/parkingfees"
	"parking/service/vehicle"
	"time"
)

type IParkingGate interface {
	CheckIn(vehicle vehicle.IVehicle) error
	CheckOut(vehicleNumber string,exitTime time.Time)  (int,error)
}

type parkingGate struct {
	parkingManager *parking.ParkingManager
	VehicleMap map[string]vehicle.IVehicle
}

func NewParkingGate(parkingManager *parking.ParkingManager) *parkingGate {
	return &parkingGate{
		parkingManager: parkingManager,
		VehicleMap: make(map[string]vehicle.IVehicle),
	}
}

func (pg *parkingGate) CheckIn(vehicle vehicle.IVehicle) error {
	vehicleNumber := vehicle.GetVehicle().Id
	if _,exists:=pg.VehicleMap[vehicleNumber];!exists {
		pg.VehicleMap[vehicleNumber] = vehicle
	}
	_, err := pg.parkingManager.AssignAvailableSpot(vehicle.GetVehicle().Id, vehicle.GetVehicle().VehicleType)
	if err != nil {
		return fmt.Errorf("failed to assign a parking spot, slot full")
	}
	return nil
}

func (pg *parkingGate) CheckOut(vehicleNumber string, exitTime time.Time) (int,error) {
	vehicle := pg.VehicleMap[vehicleNumber]
	pg.parkingManager.FreeParkingSpot(vehicle.GetVehicle().Id)
	feeStrategy, err := parkingfees.GetFeeStrategy(vehicle.GetVehicle().VehicleType)
	if err != nil {
		return 0,fmt.Errorf("error Calculating fees, err: %v", err)
	}
	parkedHours := exitTime.Hour() - vehicle.GetVehicle().EntryTime.Hour()
	calculatedFee:=feeStrategy.Calculate(parkedHours)
	return  calculatedFee, nil
}
