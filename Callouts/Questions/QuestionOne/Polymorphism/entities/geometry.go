package entities

type Geometry interface{
	Area() float32
	Perimeter() float32
	Type() string
}