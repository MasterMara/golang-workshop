package main

import "fmt"

func main() {

	myBufChan := make(chan int, 5)

	go func(myChan chan int) {

		for {

			msg := <-myChan
			fmt.Println(msg)
		}

	}(myBufChan)

	myBufChan <- 1 //Does not block because unbufferredChan
	myBufChan <- 2
	myBufChan <- 3
	myBufChan <- 4
	myBufChan <- 5
	myBufChan <- 6
	myBufChan <- 7

}

/*Some Notes By Mustafa KARACABEY:
//non-deterministic, not all values printing to output
//main goroutine stop without waiting goroutine to print all values
*/
