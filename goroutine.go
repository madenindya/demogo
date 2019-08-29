package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello world goroutine")
}
func printA() {
	for i := 0; i < 5; i++ {
		fmt.Println("a")
		time.Sleep(1 * time.Second)
	}
}
func printB() {
	for i := 0; i < 5; i++ {
		fmt.Println("b")
		time.Sleep(2 * time.Second)
	}
}

// func main() {
// 	go hello()
// 	go printA()
// 	go printB()
// 	time.Sleep(10 * time.Second)
// 	fmt.Println("main function")
// }

// goroutine + channel
func printAChan(fin chan bool) {
	for i := 0; i < 5; i++ {
		fmt.Println("a")
		time.Sleep(1 * time.Second)
	}
	fin <- true
}
func printBChan(fin chan bool) {
	for i := 0; i < 5; i++ {
		fmt.Println("b")
		time.Sleep(2 * time.Second)
	}
	fin <- true
}

// func main() {
// 	aChan := make(chan bool)
// 	bChan := make(chan bool)
// 	go printAChan(aChan)
// 	go printBChan(bChan)
// 	<-aChan
// 	<-bChan
// 	fmt.Println("main function")
// }

// goroutine + channel, a little more
func isPrime(number int, prime chan bool) {
	for i := 2; i < number/2; i++ {
		if number%i == 0 {
			prime <- false
			return
		}
	}
	prime <- true
}
func closeSqrt(number int, sqrt chan int) {
	for i := 1; i < number/2; i++ {
		if (i * i) > number {
			sqrt <- (i - 1)
			return
		}
	}
	sqrt <- 0
}

// func main() {
// 	number := 1237
// 	primechan := make(chan bool)
// 	sqrtchan := make(chan int)
// 	go isPrime(number, primechan)
// 	go closeSqrt(number, sqrtchan)

// 	boolPrime, numSqrt := <-primechan, <-sqrtchan
// 	fmt.Println("number:", number, "isPrime:", boolPrime, "closeRoundSqrt", numSqrt)
// }

// goroutine + channel, final!!
func zombie() {
	fmt.Println("Entering zombie state")
	// for {
	// }
}

// func main() {
// 	forever := make(chan bool)
// 	go zombie()
// 	<-forever
// }

// Quiz! Why error? and how to fix?
