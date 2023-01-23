package main

import (
	"fmt"
)

func main() {

	//channel initialization
	unbufferedChan := make(chan int)

	//channel declaration it will create a nil channel means not useful
	var unbufferedChan2 chan int

	fmt.Println(unbufferedChan)
	fmt.Println(unbufferedChan2) //nil

	//only can read from channel
	go func(unbufChan <-chan int) {
		//This is blocking waits until data comes
		value := <-unbufChan
		fmt.Println(value)

		//unbufChan <- 5 this is not work here
	}(unbufferedChan)

	unbufferedChan <- 1
}

/*Some Notes By Mustafa KARACABEY:

1-) Read From Channel is Blocking Waits untils data comes !!!
2-) Unbuffered Channel : Every Operations write and read is blocking.
3-) bufferred Channel : Verdiğin size a kadar blocking olmuyor ondan sonraki kısım blocking oluyor


*/
