package main

import "fmt"

//takeSquare
//showResult

func writeNumbers(myChan chan<- int) { //you can only send data Unidirectional  channel

	for i := 0; i < 10; i++ {
		myChan <- i
	}
	close(myChan) // Investiage It
}

func takeSquare(mychan <-chan int, mySqrChan chan<- int) {

	for number := range mychan {
		mySqrChan <- number * number
	}
	close(mySqrChan)
}

func showSquareResult(mySquareChan <-chan int) { //you can only read data chSquare Unidirectional channel
	for number := range mySquareChan {
		fmt.Println(number)
	}
}

func main() {

	myChan := make(chan int)
	mySquareChan := make(chan int)

	go writeNumbers(myChan)
	go takeSquare(myChan, mySquareChan)

	showSquareResult(mySquareChan)
}

/* some notes by Mustafa Karacabey

1-) bidirectional-Channel can transform Unidirectional-Channel but other direction is not possible


*/
