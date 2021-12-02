package main

import (
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
	//// Create a Channel
	c := make(chan string, 2)
	c <- "hellow"
	c <- "Raghav"

	message := <-c
	fmt.Println("mesage ", message)
	message = <- c
	fmt.Println("mesage ", message)
}

/*
func count(thing interface{}, c chan interface{}) {
	for i := 0; i < 5; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}

	close(c)
} */
