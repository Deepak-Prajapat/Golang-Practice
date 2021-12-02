package entities

type Square struct {
	Name string
	A    float32
}

func (square Square) Area() float32 {
	//Code
	return square.A * square.A
}

func (square Square) Perimeter() float32 {
	//Code
	return square.A * 4
}

func (square Square) Type() string {
	//Code
	return square.Name
}
