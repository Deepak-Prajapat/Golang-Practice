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
	fmt.Println("We are starting out program")

	developerDetails := Developer{
		memberDetails: StaffMember{
			Name: "Deepak Prajapati",
			ID:   "10",
		},
		Skills:  []string{"java", "python"},
		Package: 50000,
	}
	fmt.Println("member", developerDetails)
	fmt.Println("", developerDetails.memberDetails)
	fmt.Scanln()
}
