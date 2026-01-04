package interfaces

import (
	"fmt"
)

type Vehicle interface {
	Start() string
	Stop() string
}

type Car struct{}

func (c Car) Start() string {
	return "Car started at 5-AM"
}
func (c Car) Stop() string {
	return "Car stoped at 8-PM"
}

type Bike struct{}

func (b Bike) Start() string {
	return "Bike started at Hyderabad"
}
func (b Bike) Stop() string {
	return "Bike stopped at Secunderabad"
}

type Truck struct{}

func (t Truck) Start() string {
	return "Truck started with an Ignition"
}
func (t Truck) Stop() string {
	return "Truck stopped and parked"
}

func Problem() {
	a := Car{}
	b := Bike{}
	c := Truck{}

	vehicles := []Vehicle{a, b, c}

	for _, val := range vehicles {
		fmt.Println(val.Start())
		fmt.Println(val.Stop())
		fmt.Println("----")
	}
}
