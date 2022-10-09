package elon

import (
	"strconv"
)

// TODO: define the 'Drive()' method
func (car *Car) Drive() {
	if (car.battery - car.batteryDrain) >= 0 {
		car.battery -= car.batteryDrain
		car.distance += car.speed
	}
}

// TODO: define the 'DisplayDistance() string' method
func (car *Car) DisplayDistance() string {
	return "Driven " + strconv.Itoa(car.distance) + " meters"
}

// TODO: define the 'DisplayBattery() string' method
func (car *Car) DisplayBattery() string {
	return "Battery at " + strconv.Itoa(car.battery) + "%"
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
func (car *Car) CanFinish(trackDistance int) bool {
	iteration := int(trackDistance / car.speed)

	for i := 0; i < iteration; i++ {
		if car.battery-car.batteryDrain < 0 {
			return false
		}
		car.Drive()
	}

	return true
}
