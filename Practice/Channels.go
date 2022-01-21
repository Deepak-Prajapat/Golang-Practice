package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"
)

//// This init will clear terminal's previous output on before main loads
func init() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Fatalln("err", err)
	}
}

//Questions: Find all prime numbers from a slice, then calculate friend number of each prime number
//Then calculate max from the friend numbers slice
//then calculate factorial of maximum friend number and return it to main

func main() {
	numbers := []int{7, 8, 9, 4, 5, 6, 7, 10}

	channel := make(chan []int)           //This will get primeNumbers from findPrimes() and release them into calculateFriendNumbers()
	friendNumbersChan := make(chan []int) // This will get friends numbers from calculateFriendNumbers() and release them into calculateMax()
	factorialChan := make(chan int)       // this will get max number from calculateMax() and release into calculateFact()

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		findPrimes(numbers, channel)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		calculateFriendNumbers(channel, friendNumbersChan)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		calculateMax(friendNumbersChan, factorialChan)
	}()

	factorial := 0
	wg.Add(1)
	go func() {
		defer wg.Done()
		factorial = calculateFact(factorialChan)
	}()

	wg.Wait()
	fmt.Println("factorial ", factorial)

}

func calculateFact(calculateFact chan int) int {
	fact := 1
	for i := <-calculateFact; i > 0; i-- {
		fact *= i
	}
	return fact
}

func findPrimes(numbers []int, channel chan []int) {
	var primeNumbers []int
	for _, v := range numbers {
		isPrime := true
		for i := 2; i <= v/2; i++ {
			if v%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primeNumbers = append(primeNumbers, v)
		}
	}
	channel <- primeNumbers
}

func calculateFriendNumbers(channel chan []int, friendNumbersChan chan []int) {
	var friendNumbers []int
	for _, v := range <-channel {
		friendNumbers = append(friendNumbers, friendNumber(v))
	}
	friendNumbersChan <- friendNumbers
}

func friendNumber(num int) int {
	var friendCounter = 1
	temp := num
	for i := temp; i > 0; i /= 10 {
		friendCounter *= 10
	}
	return friendCounter - num
}

func calculateMax(friendNumberChan chan []int, factorialChan chan int) {
	numbers := <-friendNumberChan
	maxNum := numbers[0]
	for _, num := range numbers {
		if num > maxNum {
			maxNum = num
		}
	}
	factorialChan <- maxNum
}
