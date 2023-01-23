package main

import "fmt"

func sumOddNumber(oddChannel chan int, numbers ...int) {

	sum := 0

	for _, number := range numbers {

		sum += number
	}

	oddChannel <- sum //Send Data to Channel
}

//This Channel means biDirectional write and read
func sumEvenNumber(evenChannel chan int, numbers ...int) {

	sum := 0

	for _, number := range numbers {

		sum += number
	}

	evenChannel <- sum // Send data to Channel
}

func main() {

	oddChannel := make(chan int) //UnBufferedChannel
	evenChannel := make(chan int)

	go sumEvenNumber(oddChannel, 1, 2, 3, 4, 5)
	go sumOddNumber(evenChannel, 1, 2, 3, 4, 5)

	sumOdd, sumEven := <-oddChannel, <-evenChannel //Read Data From Channel Means Blocking Operation !!! Waits Until Channel data to come !!!

	fmt.Println("SumOdd:", sumOdd, "SumEven:", sumEven)

}

/*Some Notes By Mustafa KARACABEY:



 */
