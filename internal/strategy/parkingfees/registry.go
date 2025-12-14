package parkingfees

import (
	"fmt"
	"parking/internal/vehicle"
)

var feeStrategies = map[vehicle.VehicleType]FeeStrategy{
	vehicle.BikeType: &BikeFee{},
	vehicle.CarType:  &CarFee{},
}

func GetFeeStrategy(vehicleType vehicle.VehicleType) (FeeStrategy, error) {
	strategy, ok := feeStrategies[vehicleType]
	if !ok {
		return nil, fmt.Errorf("unsupported vehicle type: %s", vehicleType)
	}
	return strategy, nil
}
