package main

import (
	"fmt"
)

func simpleChanNameExample(c chan string) {

	//myValue := <-c //This is Blocking and this goroutine waits until data comes
	//fmt.Printf("My Value is %v\n", myValue)

}

func main() {

	fmt.Println("Main Goroutine Start")

	myNameChan := make(chan string)

	go simpleChanNameExample(myNameChan)

	myNameChan <- "mustafa" //Data goes to child goroutine and blocks stopped!!
	// This is blocking !! Written data has to be read first to resume futrhet operations so main goroutine will stopped !!!

	fmt.Println("Main Goroutine Ends!!")

}

/*Some Notes By Mustafa KARACABEY:

1-) https://www.velotio.com/engineering-blog/understanding-golang-channels
2-)Channel is a pipe that allows a goroutine to either put or read the data.
3-)Keep in mind that channels are nothing but pointers. That’s why we can pass them to goroutines, and we can easily put the data or read the data.
4-) Unbuffered channel is blocking read or write operations
5-) blocking and unblocking operations are done over goroutines by the Go scheduler
5-) Unless there's data in a channel you can’t read from it, which is why our simpleChanNameExample goroutine was blocked in the first place,
 the written data has to be read first to resume further operations. This happened in case of our main goroutine.
 6-) UnBuffered Channel :  A channel that can hold a single piece of data, which has to be consumed before pushing other data.
 7-) BufferedChannel:  The syntax is very simple. c := make(chan int,10)  the second argument in the make function is the capacity of a channel.
 8-) BufferedChannel  When the capacity is full, then that channel would get blocked so that the receiver goroutine can start consuming it.
9-)Channels are goroutine-safe.
10-)Channels can store and pass values between goroutines.
11-)Channels provide FIFO semantics.
12-)Channels cause goroutines to block and unblock, which we just learned about.
*/
