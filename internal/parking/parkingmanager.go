package parking

import (
	"fmt"
	"parking/internal/vehicle"

	"github.com/google/uuid"
)

type IParkingManager interface {
	Add(vehicleType string, floor int)
	Remove(parkingSpotId int)
	AssignAvailableSpot(vehicleId int, vehicleType string) (*parkingSpot, error)
	FreeParkingSpot(vehicleId int) error
}

type ParkingManager struct {
	parkingSpotIdVsPS        map[string]*parkingSpot
	parkingPerFloorPerType   map[vehicle.VehicleType]map[int][]string
	vehicleIdVsParkingSpotId map[string]string
}

func New() *ParkingManager {
	return &ParkingManager{
		parkingSpotIdVsPS:        make(map[string]*parkingSpot),
		parkingPerFloorPerType:   make(map[vehicle.VehicleType]map[int][]string),
		vehicleIdVsParkingSpotId: make(map[string]string),
	}
}

func (ps *ParkingManager) Add(vehicleType vehicle.VehicleType, floor int) {
	id := uuid.New().String()
	parkingSpot := &parkingSpot{
		id:          id,
		vehicleType: vehicleType,
		floor:       floor,
	}

	ps.parkingSpotIdVsPS[id] = parkingSpot
	if ps.parkingPerFloorPerType[vehicleType]==nil{
		ps.parkingPerFloorPerType[vehicleType]=make(map[int][]string)
	}
	ps.parkingPerFloorPerType[vehicleType][floor] = append(ps.parkingPerFloorPerType[vehicleType][floor], id)
}

func (ps *ParkingManager) Remove(parkingSpotId string) {
	parkingSpot := ps.parkingSpotIdVsPS[parkingSpotId]
	delete(ps.parkingSpotIdVsPS, parkingSpotId)
	parkingList := ps.parkingPerFloorPerType[parkingSpot.vehicleType][parkingSpot.floor]
	for i, el := range parkingList {
		if el == parkingSpotId {
			parkingList[i] = parkingList[len(parkingList)-1]
			parkingList = parkingList[:len(parkingList)-1]
		}
	}
}

func (ps *ParkingManager) AssignAvailableSpot(vehicleNumber string, vehicleType vehicle.VehicleType) (*parkingSpot, error) {
	availablefloors := ps.parkingPerFloorPerType[vehicleType]

	for _, val := range availablefloors {
		for _, id := range val {
			parkingSpot := ps.parkingSpotIdVsPS[id]
			if !parkingSpot.isOccupied {
				parkingSpot.isOccupied = true
				ps.vehicleIdVsParkingSpotId[vehicleNumber] = parkingSpot.id
				return parkingSpot, nil
			}
		}
	}
	return nil, fmt.Errorf("no spot available")
}

func (ps *ParkingManager) FreeParkingSpot(vehicleId string) error {
	parkingSpotId := ps.vehicleIdVsParkingSpotId[vehicleId]
	parkingSpot := ps.parkingSpotIdVsPS[parkingSpotId]
	parkingSpot.isOccupied = false
	delete(ps.vehicleIdVsParkingSpotId, vehicleId)
	return nil
}
