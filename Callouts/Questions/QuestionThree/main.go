package main

import (
	"fmt"
	"math"
	"math/rand"
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
	var primes []int64              // slice of prime numbers
	const maxPrime int64 = 10000000 // max value for primes

	start := time.Now()
	wg.Add(5)
	go func() {
		for i := 0; i < 2000; i++ {
			p := getPrime(maxPrime) // add a new prime
			primes = append(primes, p)
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 2000; i++ {
			p := getPrime(maxPrime) // add a new prime
			primes = append(primes, p)
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 2000; i++ {
			p := getPrime(maxPrime) // add a new prime
			primes = append(primes, p)
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 2000; i++ {
			p := getPrime(maxPrime) // add a new prime
			primes = append(primes, p)
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 2000; i++ {
			p := getPrime(maxPrime) // add a new prime
			primes = append(primes, p)
		}
		wg.Done()
	}()

	wg.Wait()
	end := time.Now()

	fmt.Println("Taken time:", end.Sub(start))
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
