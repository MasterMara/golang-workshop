package main

import "fmt"

func main() {

	//DEADLOCK !!!!
	myChan := make(chan string)

	myChan <- "Mustafa" //Write Data to channel and unbuffered Channel write/read operations is blocked !! wait untils readed to someone !!

	myMessageFromChannel := <-myChan // Read Data From Channel but main goroutine is blocked !!!

	fmt.Println("My Message is", myMessageFromChannel)

}

/*Some Notes By Mustafa KARACABEY:
1-) if no sharing data means, no deadlock, no race condition
2-) Sender No Receiver Sitation.
*/
