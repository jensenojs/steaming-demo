package vehicle_count_job

import (
	"streamwork/pkg/engine/operator"
	"streamwork/pkg/engine/source"
)

type carType = string

type VehicleEvent struct {
	Type carType
}

func (v *VehicleEvent) IsEvent() {}

// SensorReader is a monitor on the brigde, Track how many cars are passing by. specific to the type of the car
type SensorReader struct {
	source.Source
}

// VehicleCounter is a counter
type VehicleCounter struct {
	operator.Operator

	counter map[carType]int
}
