package main

import "fmt"

func main() {
	//channel initialization
	bufferedChan := make(chan int, 5)

	bufferedChan <- 1                     //Write
	fmt.Println("First:", <-bufferedChan) //Read
	bufferedChan <- 2
	fmt.Println("First:", <-bufferedChan)
	bufferedChan <- 3
	fmt.Println("First:", <-bufferedChan)
	bufferedChan <- 4
	fmt.Println("First:", <-bufferedChan)
	bufferedChan <- 5
	fmt.Println("First:", <-bufferedChan)

	fmt.Println("**********************")

	bufferedChan <- 10
	fmt.Println("asfmaşfmasçfma")

	fmt.Println("Len:", <-bufferedChan) //Read is always blocking
	fmt.Println("Len:", <-bufferedChan)
	fmt.Println("Len:", <-bufferedChan)
	fmt.Println("Len:", <-bufferedChan)
	fmt.Println("Len:", <-bufferedChan)
}

/* Some Notes By Mustafa Karacabey
- Read operations are blocking whether is buffered or unBuffered
*/
