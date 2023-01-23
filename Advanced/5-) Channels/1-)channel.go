package main

import "fmt"

func firtsGoroutine(c chan bool) {

	fmt.Println("I am Aother goroutine")

	c <- true //Send Data into Channel

}

func main() {

	done := make(chan bool) //UnBufferedChannel

	fmt.Printf("Chan type is %T\t Chan value is : %v", done, done) // The channel value is pointer

	go firtsGoroutine(done)

	<-done //Blocking  Read data from channel

	fmt.Println("Main Goroutine")
}

/*Some Notes By Mustafa KARACABEY:

1-) In Golang, or Go, channels are a means through which different goroutines communicate.
2-) Read data from channel is blocking operaiton
3-)  Keep in mind that channels are nothing but pointers.
*/
