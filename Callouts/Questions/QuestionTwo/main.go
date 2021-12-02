package main

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
	"time"
)

func clearTerminal() {
	//// This will clear terminal's previous output
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
func main() {
	clearTerminal()
	var wg sync.WaitGroup

	for _, value := range []int{9, 35, 27, 56, 88, 80} {

		wg.Add(1)
		go func(value int) {
			fmt.Println(process(value))
			wg.Done()
		}(value)
	}

	wg.Wait()
}

func process(v int) int {
	time.Sleep(1500 * time.Millisecond) // simulate compute time
	return 2 * v
}
