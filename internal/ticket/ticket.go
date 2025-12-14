package ticket

import (
	"parking/internal/vehicle"
	"time"

	"github.com/google/uuid"
)

func NewTicket(
	vehicle vehicle.IVehicle,
	floorLevel int,
	parkingSlotID string,
) *Ticket {
	return &Ticket{
		TicketID:      uuid.New().String(),
		VehicleNumber: vehicle.VehicleNumber(),
		VehicleType:   vehicle.VehicleType(),
		FloorLevel:    floorLevel,
		ParkingSlotID:        parkingSlotID,
		EntryTime:     time.Now(),
	}
}
