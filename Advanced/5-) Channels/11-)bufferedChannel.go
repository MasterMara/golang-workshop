package main

import "fmt"

func main() {

	myBufChan := make(chan int, 5)

	go func(myChan chan int) {

		for {
			fmt.Println("Len is : ", len(myChan))
			fmt.Println(<-myChan)
		}

	}(myBufChan)

	myBufChan <- 1
	myBufChan <- 2
	myBufChan <- 3
	myBufChan <- 4
	myBufChan <- 5
	myBufChan <- 6
	myBufChan <- 7
	myBufChan <- 8
	myBufChan <- 9

}

/*Some Notes By Mustafa KARACABEY:

1-)non-deterministic, not all values printing to output
2-)main goroutine stop without waiting goroutine to print all values
3-) Buffered Channel is tube if read data from them then len will decrease

*/
