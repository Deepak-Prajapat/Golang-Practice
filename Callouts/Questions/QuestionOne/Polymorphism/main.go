package main

import (
	"Callouts/Questions/QuestionOne/Polymorphism/entities"
	"fmt"
	"os"
	"os/exec"
)

func clearTerminal() {
	//// This will clear terminal's previous output
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	clearTerminal()
	circle1 := entities.Circle{"Circle 1", 5}
	circle2 := entities.Circle{"Circle 2", 15}

	rectangle1 := entities.Rectangle{"Recangle 1", 2, 4}
	rectangle2 := entities.Rectangle{"Recangle 2", 10, 4}

	square1 := entities.Square{"Square 1", 5}
	square2 := entities.Square{"Square 2", 15}

	fmt.Println("hello", circle1, circle2, rectangle1, rectangle2, square1, square2)
	geometries := []entities.Geometry{
		circle1,
		square1,
		rectangle1,
		square2,
		rectangle2,
		circle2,
	}

	fmt.Println("geome", geometries)

	for _, geometry := range geometries {
		fmt.Println(geometry.Type())
		fmt.Println("", geometry.Area())
		fmt.Println("", geometry.Perimeter())
		fmt.Println("--------------0000-----------------")
	}

	fmt.Scanln()
}
