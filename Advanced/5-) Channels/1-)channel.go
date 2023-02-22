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
4-) Kanallar go rutinleri birbiri ile iletişim kurabilmesi için kullanılan borular olarak düşünülebilir
5-) Unbuferred channel is used to perform sync comminication between goroutines
6-) buffered channel is used to perform async communication between goroutiners. no garanti
7-) In Buffered channel there is capacity to hold one or more values before they receive.
8-) buffered channelerada goroutine bitmesini beklemez.
9-) close(can)


*/
