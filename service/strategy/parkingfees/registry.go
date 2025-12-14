package parkingfees

import "fmt"

var feeStrategies = map[string]FeeStrategy{
	"bike": &BikeFee{},
	"car":  &CarFee{},
}

func GetFeeStrategy(vehicleType string) (FeeStrategy, error) {
	strategy, ok := feeStrategies[vehicleType]
	if !ok {
		return nil, fmt.Errorf("unsupported vehicle type: %s", vehicleType)
	}
	return strategy, nil
}
