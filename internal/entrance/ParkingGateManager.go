package parkinggate

import (
	"fmt"
	"parking/internal/parking"
	"parking/internal/strategy/parkingfees"
	"parking/internal/ticket"
	"parking/internal/vehicle"
	"time"
)

type IParkingGate interface {
	Park(vehicle vehicle.IVehicle) (*ticket.Ticket, error)
	Exit(vehicleNumber string, exitTime time.Time) (*ticket.Ticket, error)
}

type parkingGate struct {
	parkingManager      *parking.ParkingManager
	vehicleMap          map[string]vehicle.IVehicle
	ticketIdToTicketMap map[string]*ticket.Ticket
}

func NewParkingGate(parkingManager *parking.ParkingManager) *parkingGate {
	return &parkingGate{
		parkingManager: parkingManager,
		vehicleMap:     make(map[string]vehicle.IVehicle),
		ticketIdToTicketMap: make(map[string]*ticket.Ticket),
	}
}

func (pg *parkingGate) Park(vehicle vehicle.IVehicle) (*ticket.Ticket, error) {
	vehicleNumber := vehicle.VehicleNumber()
	if _, exists := pg.vehicleMap[vehicleNumber]; !exists {
		pg.vehicleMap[vehicleNumber] = vehicle
	}
	parkingSpot, err := pg.parkingManager.AssignAvailableSpot(vehicle.VehicleNumber(), vehicle.VehicleType())
	if err != nil {
		return nil, fmt.Errorf("failed to assign a parking spot, slot full")
	}
	ticket := ticket.NewTicket(vehicle, parkingSpot.GetFloor(), parkingSpot.GetId())
	pg.ticketIdToTicketMap[ticket.TicketID]=ticket
	return ticket, nil
}

func (pg *parkingGate) Exit(ticketId string, exitTime time.Time) (*ticket.Ticket, error) {
 	ticket := pg.ticketIdToTicketMap[ticketId]
	if ticket == nil {
		return nil, fmt.Errorf("invalid ticket")
	}
	pg.parkingManager.FreeParkingSpot(ticket.VehicleNumber)
	feeStrategy, err := parkingfees.GetFeeStrategy(ticket.VehicleType)
	if err != nil {
		return nil, fmt.Errorf("error Calculating fees, err: %v", err)
	}
	parkedHours := exitTime.Hour() - ticket.EntryTime.Hour()
	calculatedFee := feeStrategy.Calculate(parkedHours)
	ticket.ExitTime = exitTime
	ticket.Charges = calculatedFee
	return ticket, nil
}
