package structexp

import (
	"fmt"
	"strconv"
	"time"
)

type Engine struct {
	EngineName string
	EngineType string
	HorsePower int
}

type Car struct {
	Engine  Engine
	CarName string
	Brand   string
	Model   string
	Year    int
	Color   string
}

func (c Car) Display() string {
	return "Car Name: " + c.CarName + ", Brand: " + c.Brand + ", Model: " + c.Model +
		", Year: " + strconv.Itoa(c.Year) + ", Color: " + c.Color +
		", Engine: " + c.Engine.EngineName + ", Type: " + c.Engine.EngineType +
		", Horse Power: " + strconv.Itoa(c.Engine.HorsePower)
}

func GetCar() {
	Aengine := Engine{
		EngineName: "V8",
		EngineType: "Petrol",
		HorsePower: 450,
	}
	Car := Car{
		Engine:  Aengine,
		CarName: "Mustang",
		Brand:   "Ford",
		Model:   "GT",
		Year:    2021,
		Color:   "Red",
	}
	fmt.Println(Car.Display())
}

func CreateMultCar(E Engine, CName ...string) []Car {
	var cars []Car
	for _, name := range CName {

		cars = append(cars, Car{
			Engine:  E,
			CarName: name,
			Brand:   "DefaultBrand",
			Model:   "DefaultModel",
			Year:    time.Now().Year(),
			Color:   "DefaultColor",
		})
	}
	return cars
}
