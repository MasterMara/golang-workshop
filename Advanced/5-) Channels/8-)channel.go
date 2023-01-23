package main

import (
	"fmt"
	"time"
)

func main() { //Main Goroutines

	done := make(chan bool)

	go func() {
		fmt.Println("Series-1")
		time.Sleep(3 * time.Second)
		fmt.Println("Done Goroutines")
		done <- true //Write data to Channel, waits until reads to someone !!! if dont read just blocked !!!
	}()

	fmt.Println("Series-2")

	<-done // Read Data From Channel, This is blocking if channel has no value just deadlock stops blocked !!!
	fmt.Println("Done Whole Applicaiton")
}

/*Some Notes By Mustafa KARACABEY:
1-) we can send value from child goroutine to main goroutine with channels !!!!
*/
