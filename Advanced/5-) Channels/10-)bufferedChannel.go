package main

import "fmt"

func main() {

	myBufChan := make(chan int, 5)
	myUnBufChannel := make(chan bool)

	go func(bufChan chan int, myUnBufChan chan bool) {

		//when the bufChan closed, loop will be break. It will wait just if do not close channel!!!
		for val := range bufChan {
			fmt.Println(val)
		}
		fmt.Println("channel closed")

		myUnBufChan <- true

	}(myBufChan, myUnBufChannel)

	myBufChan <- 1
	myBufChan <- 2
	myBufChan <- 3
	myBufChan <- 4
	myBufChan <- 5
	myBufChan <- 6
	myBufChan <- 7
	myBufChan <- 8
	close(myBufChan)

	//wait goroutine to finish
	<-myUnBufChannel
	fmt.Println("main finished")

}

/*Some Notes By Mustafa KARACABEY:


 */
