package vehicle

import "time"

const (
	CAR  = "car"
	BIKE = "bike"
)

type Vehicle struct {
	Id            string
	VehicleType      string
	EntryTime     time.Time
}

