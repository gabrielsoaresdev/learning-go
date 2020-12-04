package main

import "fmt"

const usixteenbitmax float64 = 65535
const kmh_multiple float64 = 1.60934

type car struct {
	gas_pedal uint16 //min 0 max 65535
	brake_pedal uint16
	steering_wheel int16 //-32k ~ +32k	
	top_speed_kmh float64
}

func (this car) kmh() float64 {
	return float64(this.gas_pedal) * (this.top_speed_kmh/usixteenbitmax)
}

func (c car) mmh() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh/usixteenbitmax/kmh_multiple)
}

func main() {
	a_car := car {
		gas_pedal: 22341,
		brake_pedal: 0,
		steering_wheel: 12561,
		top_speed_kmh: 225.0}

	fmt.Println(a_car.gas_pedal)
	fmt.Println(a_car.gas_pedal)
	fmt.Println(a_car.kmh())
}