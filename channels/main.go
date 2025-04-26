package main

import (
	"fmt"
)

// send
// func processNum(numChan chan int) {
// 	for num := range numChan {
// 		fmt.Println("processing number", num)
// 		time.Sleep(time.Second)
// 	}
// }

// receive
// func sum(result chan int, num1 int, num2 int) {
// 	numResult := num1 + num2
// 	result <- numResult
// }

// func task(done chan bool) {
// 	defer func() {
// 		done <- true
// 	}()
// 	fmt.Println("processing...")
// }

// func emailSender(emailChan chan string, done chan bool) {
// 	defer func() { done <- true }()
// 	for email := range emailChan {
// 		fmt.Println("sending email to ", email)
// 		time.Sleep(time.Second)
// 	}
// }

func main() {
	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()

	go func() {
		chan2 <- "Pong"
	}()

	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("fron chan1", chan1Val)
		case chan2Val := <-chan2:
			fmt.Println("chan 2 value", chan2Val)
		}
	}

	// emailChan := make(chan string, 100)
	// done := make(chan bool)

	// go emailSender(emailChan, done)

	// for i := 0; i < 5; i++ {
	// 	emailChan <- fmt.Sprintf("%d@gmail.com", i)
	// }
	// fmt.Println("done sending")

	// close(emailChan)
	// <-done
	// done := make(chan bool)
	// go task(done)

	// <-done
	// result := make(chan int)

	// go sum(result, 4, 5)
	// res := <-result
	// fmt.Println(res)
	// numChan := make(chan int)

	// go processNum(numChan)

	// for {
	// 	numChan <- rand.Intn(100)
	// }
	// messageChan := make(chan string)
	// messageChan <- "ping"
	// msg := <-messageChan
	// fmt.Println(msg)
}
