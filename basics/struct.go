package main

import "fmt"

const uint16max float64 = 65535
const kmh_multiple float64 = 1.60934

type car struct {
	gasPedal      uint16 // 0 ~ 65535
	brakePedal    uint16
	steeringWheel int16 // -32768 ~ 32767
	topSpeedKmh   float64
	Price         uint32
}

// Value Receiver
// c - car struct
// Add kmh() into car struct!
// (Like extension method in Kotlin)
func (c car) kmh() float64 {
	return float64(c.gasPedal) * c.topSpeedKmh / uint16max
}

// Pointer Receiver
func (c *car) newTopSpeed(speed float64) {
	c.topSpeedKmh = speed
}

func main() {
	newCar := car{
		gasPedal:      24000,
		brakePedal:    0,
		steeringWheel: 14582,
		topSpeedKmh:   140.0,
		Price:         20000}

	// simpleCar := car{2323, 0, 14023, 500.0, 30000}

	fmt.Println(newCar.Price)
	fmt.Println(newCar.kmh())
}
