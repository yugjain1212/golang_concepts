// package main

// import (
// 	"crypto/rand"
// 	"fmt"
// 	"time"
// )

// func processnum(message chan int) {
// 	fmt.Println("message", <-message)
// 	for num := range message {
// 		fmt.Println("processing number", num)
// 		time.Sleep(time.Second)
// 	}

// }
// func main() {
// 	message := make(chan int)
// 	go processnum(message)
// 	message <- 5
// 	for {
// 		message <- rand.Intn(100)
// 	}
// 	time.Sleep(time.Second * 2)

//		// message <- "hello"
//		// msg := <-message
//		// fmt.Println(msg)
//	}
package main

import (
	"fmt"
)

// func processnum(message chan int) {
// 	fmt.Println("message", <-message)
// 	for num := range message {
// 		fmt.Println("processing number", num)
// 		time.Sleep(time.Second)
// 	}

// }
//
//	func sum(result chan int, nun1 int, nun2 int) {
//		result <- nun1 + nun2
//	}
// func task(done chan bool) {
// 	defer func() {
// 		done <- true
// 	}()
    
// 	fmt.Println("processing....")

//  }
// func emailsend(email chan string, done chan bool) {
//   defer func() {
// 		done <- true
// 	}()
// 	for email := range email {
// 		fmt.Println("sending email to ", email)
// 		time.Sleep(time.Second) // sleep for 1 second
//	}
//	}
func main() {
	chan1 := make(chan int)
	chan2 := make(chan string)
	go func() {
		chan1 <- 10
	}()
	go func() {
		chan2 <- "hello"
	}()
	for i := 0; i < 2; i++ {
		select {
		case x := <-chan1:
			fmt.Println(x)
		case y := <-chan2:
			fmt.Println(y)
		}
	}

	// email := make(chan string, 100)
	// done := make(chan bool)
	// go emailsend(email, done)
	// for i := 0; i < 100; i++ {
	// 	email <- fmt.Sprintf("email %d", i)
	// }
	// fmt.Println("email sent done")
	// //important to close the channel
	// close(email)
	// <-done

	// email <- "1@example.com"
	// email <- "2@example.com"
	// fmt.Println(<-email)
	// fmt.Println(<-email)
	// res := make(chan bool)
	// go task(res)
	// <-res //block until task is done

	// result := make(chan int)
	// go sum(result, 4, 5)
	// re := <-result
	// fmt.Println(re)

	// message := make(chan int)
	// go processnum(message)
	// message <- 5
	// for {
	// 	message <- rand.Intn(100)
	// }
}
