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
	fmt.Println("helloworld")
	var myString string = "Hello Chutiye"
	fmt.Println( myString)
	myString = 

	fmt.Print("Enter String = ")
	varName, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	varName = varName[0: len(varName) - 2]

}
