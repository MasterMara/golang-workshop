package main

import "fmt"

func main() {

	myBufChannel := make(chan string, 2)

	myBufChannel <- "Mustafa" //Does not blocking because this is a buffered channel we have a capacity to cover data
	myBufChannel <- "Kadir"

	msg := <-myBufChannel

	fmt.Println(msg)

	msg = <-myBufChannel
	fmt.Println(msg)

}

/*Some Notes By Mustafa KARACABEY:
1-) There is no deadlock beacuse we have capacitiy in buffered channel
2-) it will read first value in channel fifo
3-) //Kanalları bir havuz olarak düşünebilirsin burada buffered channel create ettik ve içerisine datayı attık bu bloklamaya neden olmadı
	// Eğer capaciteyi aşarsan patlarsın :) gene deadlock
*/
