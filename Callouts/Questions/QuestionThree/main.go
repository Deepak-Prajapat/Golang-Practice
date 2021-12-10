package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

// returns true if number is prime func
func isPrime(v int64) bool {
	sq := int64(math.Sqrt(float64(v))) + 1

	var i int64
	for i = 2; i < sq; i++ {
		if v%i == 0 {
			return false
		}
	}
	return true
}

// get a random prime number between 1 and maxP
func getPrime(maxP int64) int64 {
	var i int64
	for i = 0; i < maxP; i++ {
		n := rand.Int63n(maxP)
		if isPrime(n) {
			return n
		}
	}
	return 1 // just in case
}
func main() {

	channel := make(chan int64)
	const maxPrime int64 = 10000000 // max value for primes
	start := time.Now()
	for i := 0; i < 5; i++ {
		go func(maxPrime int64, c chan int64) {
			for i := 0; i < 2000; i++ {
				c <- getPrime(maxPrime)
			}
		}(maxPrime, channel)
	}
	var primes []int64
	for i := 0; i < 10000; i++ {
		primes = append(primes, <-channel)
	}
	end := time.Now()

	fmt.Println("End of program.", end.Sub(start))
	fmt.Println("len", len(primes))
}

//// This init will clear terminal's previous output on before main loads
func init() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
