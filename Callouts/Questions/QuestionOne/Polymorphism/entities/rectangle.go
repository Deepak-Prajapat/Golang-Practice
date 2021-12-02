package entities

type Rectangle struct {
	Name string
	A    float32
	B    float32
}

func (rectangle Rectangle) Area() float32 {
	//Code
	return rectangle.A * rectangle.B
}

func (rectangle Rectangle) Perimeter() float32 {
	//Code
	return (rectangle.A + rectangle.B) * 2
}

func (rectangle Rectangle) Type() string {
	//Code
	return rectangle.Name
}
