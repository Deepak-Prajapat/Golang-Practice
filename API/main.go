package main

import (
	"os"
	"os/exec"
)

func main() {
	//// This will clear terminal's previous output
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()

	//// Work start from here
	DataMigration()
	handlerRouting()

}
