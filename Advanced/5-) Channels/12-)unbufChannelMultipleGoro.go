package main

import "fmt"

func main() {

	myChan := make(chan int)

	//Writer Goroutine
	go func(myChan chan int) {
		myChan <- 1
	}(myChan)

	//Read Goroutine
	go func(myChan chan int) {
		value := <-myChan
		fmt.Println(value)
	}(myChan)

	fmt.Println("Finished Main Goroutine")
}

/* some Notes by Mustafa Karacabey
Output is non-deterministic. Scheduler probably will not have time to schedule goroutines.
So we will not see channel value in the output
*/
