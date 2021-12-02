package entities

import "math"

type Circle struct {
	Name string
	R    float32
}

func (circle Circle) Area() float32 {
	//Code
	return math.Pi * circle.R * circle.R
}

func (circle Circle) Perimeter() float32 {
	//Code
	return 2 * math.Pi * circle.R
}

func (circle Circle) Type() string {
	//Code
	return circle.Name
}

type Dog struct {
	Name string
}

func (dog Dog) Speak() string {
	return "Bhau Bhau"
}

func (dog Dog) AnimalName() string {
	return dog.Name
}
