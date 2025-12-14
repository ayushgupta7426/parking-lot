package ticket

import (
	"parking/internal/vehicle"
	"time"
)

type Ticket struct {
	TicketID      string
	VehicleNumber string
	VehicleType   vehicle.VehicleType
	FloorLevel    int
	ParkingSlotID        string
	EntryTime     time.Time
	ExitTime time.Time
	Charges int
}